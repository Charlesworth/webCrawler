package main

import "net/url"

func stringIsValidURL(urlString string) bool {
	url, err := url.Parse(urlString)
	if err != nil {
		return false
	}

	if !validHTTPOrHTTPS(url) {
		return false
	}

	return true
}

func validHTTPOrHTTPS(url *url.URL) bool {
	return ((url.Scheme == "http") || (url.Scheme == "https"))
}

func hasSchemeAndHost(queryURL string) bool {
	parsedURL, err := url.Parse(queryURL)
	if err != nil {
		return false
	}
	if parsedURL.Scheme == "" || parsedURL.Host == "" {
		return false
	}
	return true
}

func isSameDomain(queryURL string, baseURL string) bool {
	parsedQueryURL, err := url.Parse(queryURL)
	if err != nil {
		return false
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		return false
	}

	return (parsedQueryURL.Scheme == parsedBaseURL.Scheme) && (parsedQueryURL.Host == parsedBaseURL.Host)
}

func appendSchemeAndHost(href string, base string) string {
	uri, err := url.Parse(href)
	if err != nil {
		return href
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		return href
	}

	temp := baseURL.ResolveReference(uri)
	return temp.String()
}
