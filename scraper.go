package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/emmerjo/goprojectrss/internal/database"
)

func startScraping(
	db *database.Queries,
	concurrency int,
	timeBetweenRequest time.Duration,
) {
	log.Printf("Scraping op %v goroutines iedere %s duratie", concurrency, timeBetweenRequest)
	ticker := time.NewTicker(timeBetweenRequest)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(
			context.Background(),
			int32(concurrency),
		)
		if err != nil {
			log.Println("error bij fetchen feeds", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)

			go scrapeFeed(db, wg, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(db *database.Queries, wg *sync.WaitGroup, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Println("Error bij het markeren van feed als fetched", err)
		return
	}

	rssFeed, err := urlToFeed(feed.Url)
	if err != nil {
		log.Println("Error bij fetchen van feed: ", err)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		log.Println("Post gevonden", item.Title, "op feed", feed.Name)
	}
	log.Printf("Feed %s verzameld, %v posts gevonden", feed.Name, len(rssFeed.Channel.Item))
}
