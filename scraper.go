package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/eduvedras/Blog-Aggregator/internal/database"
)

func startScraping(db *database.Queries, concurrency int, timeBetweenRequest time.Duration) {
	log.Printf("Collecting feeds every %s on %v goroutines...", timeBetweenRequest, concurrency)
	ticker := time.NewTicker(timeBetweenRequest)

	for ; ; <-ticker.C{
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(concurrency))
		if err != nil{
			log.Println("Couldn't get next feeds to fetch")
			continue
		}
		log.Printf("Found %v feeds to fetch!", len(feeds))

		wg := &sync.WaitGroup{}
		for _, feed := range feeds{
			wg.Add(1)
			go scrapeFeed(db, feed, wg)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, feed database.Feed, wg *sync.WaitGroup){
	defer wg.Done()

	_ , err := db.MarkFeedFetch(context.Background(), feed.ID)
	if err != nil{
		log.Printf("Couldn't mark feed %s fetched: %v", feed.Name, err)
		return
	}

	rssFeed, err := feedFetch(feed.Url)
	if err != nil{
		log.Printf("Couldn't collect feed %s: %v", feed.Name, err)
		return
	}

	for _, item := range rssFeed.Channel.Item{
		log.Println("Found post", item.Title)
	}

	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}

type RSSItem struct{
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

type RSSFeed struct{
	Channel struct{
		Title string `xml:"title"`
		Link string `xml:"link"`
		Description string `xml:"description"`
		Language string `xml:"language"`
		Item []RSSItem `xml:"item"`
	} `xml:"channel"`
}

func feedFetch(feedURL string) (*RSSFeed, error){
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(feedURL)
	if err != nil{
		return nil, err
	}

	dat, err := io.ReadAll(resp.Body)
	if err != nil{
		return nil, err
	}
	
	var rssFeed RSSFeed
	err = xml.Unmarshal([]byte(dat), &rssFeed)
	if err != nil{
		return nil, err
	}

	fmt.Println(rssFeed)
	return &rssFeed, nil
}