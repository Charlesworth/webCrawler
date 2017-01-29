package main

import (
	"errors"
	"net/url"
)

func stringToValidURL(urlString string) error {
	url, err := url.Parse(urlString)
	if err != nil {
		return err
	}

	if !validHTTPOrHTTPS(url) {
		return errors.New("require a valid http or https address")
	}

	return nil
}

func validHTTPOrHTTPS(url *url.URL) bool {
	return ((url.Scheme == "http") || (url.Scheme == "https"))
}
