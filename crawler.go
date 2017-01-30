package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oleiade/lane"
)

var crawledURLs = make(map[string]bool)
var crawlQueue = lane.NewQueue()

func crawl(startingURL string) {
	crawlQueue.Enqueue(startingURL)
	fmt.Println("[")

	for {
		if crawlQueue.Empty() {
			break
		}

		currentURL := crawlQueue.Dequeue().(string)
		resp, err := http.Get(currentURL)
		if err != nil {
			continue
		}

		links := getLinks(resp)
		intraDomainLinks := sortSameDomain(links, startingURL)

		assets := getAssets(resp)

		for _, link := range intraDomainLinks {
			_, present := crawledURLs[link]
			if !present {
				crawlQueue.Enqueue(link)
				crawledURLs[link] = true
			}
		}

		printJSON(currentURL, assets)
	}

	fmt.Println("]")
}

type page struct {
	url    string
	assets []string
}

func printJSON(url string, assets []string) {
	page := page{
		url,
		assets,
	}

	res, err := json.Marshal(page)
	if err != nil {
		return
	}

	fmt.Println(string(res), ",")
}
