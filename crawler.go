package main

import (
	"errors"
	"net/http"

	"github.com/charlesworth/webCrawler/jsonPrinter"

	"github.com/oleiade/lane"
)

var crawledURLs map[string]bool
var crawlQueue *lane.Queue

// crawl takes a starting URL and crawls through all hyperlinks, printing results to std out.
func crawl(startingURL string) error {
	// initialise the crawledURL map and que
	crawledURLs = make(map[string]bool)
	crawlQueue = lane.NewQueue()

	// enque the startingURL and being the crawling loop
	crawlQueue.Enqueue(startingURL)

	for {
		// if there are no more links to crawl in the que, break
		if crawlQueue.Empty() {
			break
		}

		// get the next link on the queue
		currentURL := crawlQueue.Dequeue().(string)

		// http get the currentURL
		resp, err := http.Get(currentURL)
		if err != nil {
			// if there is an error and it's the starting URL, return an error, else continue
			if currentURL == startingURL {
				return errors.New("Error reaching starting URL:" + err.Error())
			}
			continue
		}

		// get the links and assets from the response body
		links, assets := getLinksAndAssets(resp)
		resp.Body.Close()

		// sort the links to return all internal links in their absolute form
		intraDomainLinks := sortSameDomain(links, startingURL)

		// add those links to the crawl queue
		addLinksToCrawl(intraDomainLinks)

		// print the current page's URL and assets in JSON format
		jsonPrinter.PrintPage(currentURL, assets)
	}

	// end the JSON printing by closing the array square bracket
	jsonPrinter.End()
	return nil
}

// for each link, check if its in the crawled map, if not, add it to both the map and queue
func addLinksToCrawl(intraDomainLinks []string) {
	for _, link := range intraDomainLinks {
		_, present := crawledURLs[link]
		if !present {
			crawlQueue.Enqueue(link)
			crawledURLs[link] = true
		}
	}
}
