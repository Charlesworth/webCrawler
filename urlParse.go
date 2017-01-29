package main

import (
	"errors"
	"net/url"
)

func stringToValidURL(urlString string) (*url.URL, error) {
	url, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	if !validHTTPOrHTTPS(url) {
		return nil, errors.New("require a valid http or https address")
	}

	return url, nil
}

func validHTTPOrHTTPS(url *url.URL) bool {
	return ((url.Scheme == "http") || (url.Scheme == "https"))
}
