package main

import (
	"net/http"

	"golang.org/x/net/html"
)

// getLinksAndAssets returns slice of links and a slice of assets contained in the supplied response body
func getLinksAndAssets(response *http.Response) (links []string, assets []string) {
	htmlTokenizer := html.NewTokenizer(response.Body)
	links = []string{}
	assets = []string{}

	// for each token in the response body, until we reach the end of the body (where tokenType == html.ErrorToken)
	for tokenType := htmlTokenizer.Next(); tokenType != html.ErrorToken; tokenType = htmlTokenizer.Next() {

		// if the token type is a starting tag, e.g. <a> or <img>
		if tokenType == html.StartTagToken {
			// get that token
			token := htmlTokenizer.Token()

			// if the token is type <a>
			if token.Data == "a" {
				for _, attribute := range token.Attr {
					// range over all attributes, if a href is found and is not a link to the same page, add it to links slice
					if attribute.Key == "href" && !internalLink(attribute.Key) {
						links = append(links, attribute.Val)
					}
				}
			}

			// if the token is type <img>, <script> or <source> (anything that could use an asset)
			if token.Data == "img" || token.Data == "script" || token.Data == "source" {
				// range over all attributes
				for _, attribute := range token.Attr {
					// if a href  or scr is found
					if (attribute.Key == "src") || (attribute.Key == "href") {
						// add scheme and host if not present in URL and then add to assets slice
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
