package main

import (
	"net/http"

	"golang.org/x/net/html"
)

func getLinks(response *http.Response) []string {
	htmlTokenizer := html.NewTokenizer(response.Body)
	links := []string{}

	for tokenType := htmlTokenizer.Next(); tokenType != html.ErrorToken; tokenType = htmlTokenizer.Next() {

		if tokenType == html.StartTagToken {
			token := htmlTokenizer.Token()

			if token.Data == "a" {
				for _, attribute := range token.Attr {
					if attribute.Key == "href" {
						links = append(links, attribute.Val)
					}
				}
			}

		}

	}

	return links
}

func getAssets(response *http.Response) []string {
	htmlTokenizer := html.NewTokenizer(response.Body)
	links := []string{}

	for tokenType := htmlTokenizer.Next(); tokenType != html.ErrorToken; tokenType = htmlTokenizer.Next() {

		if tokenType == html.StartTagToken {
			token := htmlTokenizer.Token()

			if token.Data == "img" || token.Data == "script" || token.Data == "source" {
				for _, attribute := range token.Attr {
					if (attribute.Key == "src") || (attribute.Key == "href") {
						if !hasSchemeAndHost(attribute.Val) {
							newURL := appendSchemeAndHost(attribute.Val, response.Request.URL.String())
							links = append(links, newURL)
						} else {
							links = append(links, attribute.Val)
						}
					}
				}
			}

		}

	}

	return links
}
