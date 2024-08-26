package main

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/exec"
	"testing"
)

const host = "localhost:8080/v1"

func TestCreateAndGetUser(t *testing.T) {
	cases := []struct {
		name     string
		feedName string
		url      string
	}{
		{
			name:     "Jose",
			feedName: "BBC News (World)",
			url:      "http://feeds.bbci.co.uk/news/world/rss.xml",
		},
		{
			name:     "Maria",
			feedName: "ESPN (Top Headlines)",
			url:      "http://www.espn.com/espn/rss/news",
		},
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	connString := os.Getenv("CONN_STRING")

	//Delete all records in database
	deleteCmd := exec.Command("goose", "postgres", connString, "down-to", "0")
	deleteCmd.Dir = "./sql/schema/"
	err = deleteCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	migrateUpCmd := exec.Command("goose", "postgres", connString, "up")
	migrateUpCmd.Dir = "./sql/schema/"
	err = migrateUpCmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	for i, name := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			//User handlers tests
			createUserCmd := exec.Command("curl", "-X", "POST", "-d", fmt.Sprintf("{\"name\": \"%s\"}", name.name), fmt.Sprintf("%v/users", host))
			resCreateUser := User{}
			err := cmdOutput(createUserCmd, &resCreateUser)
			if err != nil {
				t.Error(err)
				return
			}

			getUserCmd := exec.Command("curl", "-X", "GET", fmt.Sprintf("%v/users", host), "-H", fmt.Sprintf("Authorization: ApiKey %s", resCreateUser.ApiKey))
			resGetUser := User{}
			err = cmdOutput(getUserCmd, &resGetUser)
			if err != nil {
				t.Error(err)
				return
			}

			if resCreateUser != resGetUser {
				t.Errorf("User from create is %v and user from get is %v", resCreateUser, resGetUser)
				return
			}

			//Feed handlers tests
			createFeedCmd := exec.Command("curl", "-X", "POST", "-d", fmt.Sprintf("{\"name\": \"%s\", \"url\": \"%s\"}", name.feedName, name.url), fmt.Sprintf("%v/feeds", host), "-H", fmt.Sprintf("Authorization: ApiKey %s", resCreateUser.ApiKey))
			resCreateFeed := struct {
				Feed       Feed       `json:"feed"`
				FeedFollow FeedFollow `json:"feed_follow"`
			}{}
			err = cmdOutput(createFeedCmd, &resCreateFeed)
			if err != nil {
				t.Error(err)
				return
			}

			getFeedsCmd := exec.Command("curl", "-X", "GET", fmt.Sprintf("%v/feeds", host))
			resGetFeeds := []Feed{}
			err = cmdOutput(getFeedsCmd, &resGetFeeds)
			if err != nil {
				t.Error(err)
				return
			}

			if !contains(resGetFeeds, resCreateFeed.Feed) {
				t.Errorf("Feed received from create %v is not inside the feeds from the get command %v", resCreateFeed.Feed, resGetFeeds)
				return
			}

			//Feed follows integration tests
			/*getFeedFollowsCmd := exec.Command("curl", "-X", "GET", fmt.Sprintf("%v/feed_follows", host), "-H", fmt.Sprintf("Authorization: ApiKey %s", resCreateUser.ApiKey))
			resGetFeedFollows := []FeedFollow{}
			err = cmdOutput(getFeedFollowsCmd, &resGetFeedFollows)
			if err != nil {
				t.Error(err)
				return
			}

			if !contains(resGetFeedFollows, resCreateFeed.FeedFollow) {
				t.Errorf("Feed follow %v from create feed command is not inside slice of feed follows from get feed follows command %v", resCreateFeed.FeedFollow, resGetFeedFollows)
				return
			}*/
		})
	}
}

func cmdOutput(cmd *exec.Cmd, payload any) error {
	stdOut, err := cmd.Output()
	if err != nil {
		return err
	}

	err = json.Unmarshal(stdOut, payload)
	if err != nil {
		return err
	}
	return nil
}

func contains[T comparable](s []T, f T) bool {
	for _, el := range s {
		if el == f {
			return true
		}
	}
	return false
}
