package main

func sortSameDomain(urlSlice []string, domain string) []string {
	sortedURLs := []string{}
	if len(urlSlice) == 0 {
		return sortedURLs
	}

	for _, url := range urlSlice {
		if hasSchemeAndHost(url) {
			if isSameDomain(url, domain) {
				sortedURLs = append(sortedURLs, url)
			}
		} else {
			url = appendSchemeAndHost(url, domain)
			sortedURLs = append(sortedURLs, url)
		}
	}

	return sortedURLs
}
