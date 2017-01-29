package main

import (
	"net/url"
	"testing"
)

func TestStringIsValidURL(t *testing.T) {
	HTTPURLString := "http://google.com"
	err := stringIsValidURL(HTTPURLString)
	if err != nil {
		t.Error("stringIsValidURL returned an error on a valid HTTP string:", err)
	}

	HTTPSURLString := "https://google.com"
	err = stringIsValidURL(HTTPSURLString)
	if err != nil {
		t.Error("stringIsValidURL returned an error on a valid HTTPS string:", err)
	}

	noSchemeURLString := "google.com"
	err = stringIsValidURL(noSchemeURLString)
	if err == nil {
		t.Error("stringIsValidURL returned no error on a string with no scheme")
	}

	incorrectSchemeURLString := "postgres://google.com"
	err = stringIsValidURL(incorrectSchemeURLString)
	if err == nil {
		t.Error("stringIsValidURL returned no error on a string with a non HTTP/HTTPS scheme")
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
