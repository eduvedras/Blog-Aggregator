package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"testing"
)

const host = "localhost:8080/v1"

func TestCreateUser(t *testing.T) {
	cases := []struct {
		name string
	}{
		{
			name: "Jose",
		},
		{
			name: "Maria",
		},
	}

	for i, name := range cases {
		t.Run(fmt.Sprintf("Test case %v", i), func(t *testing.T) {
			cmd := exec.Command("curl", "-X", "POST", "-d", fmt.Sprintf("{\"name\": \"%s\"}", name), fmt.Sprintf("%v/users", host))
			stdOut, err := cmd.Output()
			if err != nil {
				t.Error(err)
				return
			}

			resCreateUser := User{}

			err = json.Unmarshal(stdOut, &resCreateUser)
			if err != nil {
				t.Error(err)
				return
			}

			getCmd := exec.Command("curl", "-X", "GET", fmt.Sprintf("%v/users", host), "-H", fmt.Sprintf("Authorization: ApiKey %s", resCreateUser.ApiKey))

			stdOut, err = getCmd.Output()
			if err != nil {
				t.Error(err)
				return
			}
			resGetUser := User{}

			err = json.Unmarshal(stdOut, &resGetUser)
			if err != nil {
				t.Error(err)
				return
			}

			if resCreateUser != resGetUser {
				t.Errorf(fmt.Sprintf("User from create is %v and user from get is %v", resCreateUser, resGetUser))
				return
			}
		})
	}
}
