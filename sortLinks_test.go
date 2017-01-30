package main

import (
	"strings"
	"testing"
)

func TestSortSameDomain(t *testing.T) {
	domain := "https://www.google.com"

	allDifferentDomains := []string{"https://twitter.com/google", "mailto://www.google.com/test", "http://www.google.com"}
	sortedURLs := sortSameDomain(allDifferentDomains, domain)
	if len(sortedURLs) != 0 {
		t.Error("sortSameDomain returned urls from a different domain")
	}

	allSameDomains := []string{"https://www.google.com", "https://www.google.com/test/", "/test"}
	sortedURLs = sortSameDomain(allSameDomains, domain)
	if len(sortedURLs) != 3 {
		t.Error("sortSameDomain did not return urls from the same domain")
	}

	mixedSamilarityDomains := []string{"https://www.google.com", "https://www.google.com/test/", "/test", "https://twitter.com/google", "mailto://www.google.com/test", "http://www.google.com"}
	sortedURLs = sortSameDomain(mixedSamilarityDomains, domain)
	if len(sortedURLs) != 3 {
		t.Error("sortSameDomain returned incorrect amount of urls from a mixed list of domains")
	}
	for i := 0; i < 3; i++ {
		if !strings.Contains(sortedURLs[i], allSameDomains[i]) {
			t.Error("returned string mismatch, ", sortedURLs[i], " should contain ", allSameDomains[i])
		}
	}
}
