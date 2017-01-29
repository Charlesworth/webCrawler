package main

import (
	"net/url"
	"testing"
)

func TestStringToValidURL(t *testing.T) {
	HTTPURLString := "http://google.com"
	_, err := stringToValidURL(HTTPURLString)
	if err != nil {
		t.Error("stringToValidURL returned an error on a valid HTTP string:", err)
	}

	HTTPSURLString := "https://google.com"
	_, err = stringToValidURL(HTTPSURLString)
	if err != nil {
		t.Error("stringToValidURL returned an error on a valid HTTPS string:", err)
	}

	noSchemeURLString := "google.com"
	_, err = stringToValidURL(noSchemeURLString)
	if err == nil {
		t.Error("stringToValidURL returned no error on a string with no scheme")
	}

	incorrectSchemeURLString := "postgres://google.com"
	_, err = stringToValidURL(incorrectSchemeURLString)
	if err == nil {
		t.Error("stringToValidURL returned no error on a string with a non HTTP/HTTPS scheme")
	}
}

func TestValidHTTPOrHTTPS(t *testing.T) {
	httpURL, _ := url.Parse("http://google.com")
	if !validHTTPOrHTTPS(httpURL) {
		t.Error("validHTTPOrHTTPS returned false on a valid http url")
	}

	httpsURL, _ := url.Parse("https://google.com")
	if !validHTTPOrHTTPS(httpsURL) {
		t.Error("validHTTPOrHTTPS returned false on a valid https url")
	}

	postgresURL, _ := url.Parse("postgres://user:pass@host.com:5432/path?k=v#f")
	if validHTTPOrHTTPS(postgresURL) {
		t.Error("validHTTPOrHTTPS returned true on a non http/https url")
	}
}
