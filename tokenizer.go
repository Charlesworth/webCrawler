package main

import (
	"net/http"

	"golang.org/x/net/html"
)

// getLinks returns a slice of strings, each of which is a hyperlink in the response body
func getLinksAndAssets(response *http.Response) (links []string, assets []string) {
	htmlTokenizer := html.NewTokenizer(response.Body)
	links = []string{}
	assets = []string{}

	for tokenType := htmlTokenizer.Next(); tokenType != html.ErrorToken; tokenType = htmlTokenizer.Next() {

		if tokenType == html.StartTagToken {
			token := htmlTokenizer.Token()

			if token.Data == "a" {
				for _, attribute := range token.Attr {
					if attribute.Key == "href" && !internalLink(attribute.Key) {
						links = append(links, attribute.Val)
					}
				}
			}

			if token.Data == "img" || token.Data == "script" || token.Data == "source" {
				for _, attribute := range token.Attr {
					if (attribute.Key == "src") || (attribute.Key == "href") {
						if !hasSchemeAndHost(attribute.Val) {
							newURL := appendSchemeAndHost(attribute.Val, response.Request.URL.String())
							assets = append(assets, newURL)
						} else {
							assets = append(assets, attribute.Val)
						}
					}
				}
			}

		}

	}

	return links, assets
}
