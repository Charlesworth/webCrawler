package main

// sortSameDomain adds the domain to non absolute urls and removes any external links
func sortSameDomain(urlSlice []string, domain string) []string {
	sortedURLs := []string{}
	if len(urlSlice) == 0 {
		return sortedURLs
	}

	for _, url := range urlSlice {
		if hasSchemeAndHost(url) {
			// if the link has scheme and host information
			if isSameDomain(url, domain) {
				// if its an internal link, add to the return slice
				sortedURLs = append(sortedURLs, url)
			}
		} else {
			// else append the scheme and host and add to the return slice
			url = appendSchemeAndHost(url, domain)
			sortedURLs = append(sortedURLs, url)
		}
	}

	return sortedURLs
}
