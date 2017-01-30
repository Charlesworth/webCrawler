package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/oleiade/lane"
)

var crawledURLs = make(map[string]bool)
var crawlQueue = lane.NewQueue()

// crawl takes a starting URL and crawls through all hyperlinks, printing results to std out.
func crawl(startingURL string) {
	crawlQueue.Enqueue(startingURL)
	fmt.Fprintln(jsonPrinter, "[")

	for {
		if crawlQueue.Empty() {
			break
		}

		currentURL := crawlQueue.Dequeue().(string)
		resp, err := http.Get(currentURL)
		if err != nil {
			continue
		}

		links, assets := getLinksAndAssets(resp)
		resp.Body.Close()

		intraDomainLinks := sortSameDomain(links, startingURL)

		for _, link := range intraDomainLinks {
			_, present := crawledURLs[link]
			if !present {
				crawlQueue.Enqueue(link)
				crawledURLs[link] = true
			}
		}

		printJSON(currentURL, assets)
		fmt.Fprintln(jsonPrinter, ",")
	}

	fmt.Fprintln(jsonPrinter, "]")
}

// Page is the struct used to marshal page URL and asset info to JSON
type Page struct {
	URL    string   `json:"url"`
	Assets []string `json:"assets"`
}

// printJSON is used print a page URL and assets to std out in JSON format
func printJSON(url string, assets []string) {
	page := Page{
		url,
		assets,
	}

	res, err := json.Marshal(page)
	if err != nil {
		return
	}

	fmt.Fprint(jsonPrinter, string(res))
}
