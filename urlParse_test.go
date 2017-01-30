package main

import (
	"net/url"
	"testing"
)

func TestStringIsValidURL(t *testing.T) {
	HTTPURLString := "http://google.com"
	if !stringIsValidURL(HTTPURLString) {
		t.Error("stringIsValidURL returned false on a valid HTTP string")
	}

	HTTPSURLString := "https://google.com"
	if !stringIsValidURL(HTTPSURLString) {
		t.Error("stringIsValidURL returned false on a valid HTTPS string")
	}

	noSchemeURLString := "google.com"
	if stringIsValidURL(noSchemeURLString) {
		t.Error("stringIsValidURL returned true on a string with no scheme")
	}

	incorrectSchemeURLString := "postgres://google.com"
	if stringIsValidURL(incorrectSchemeURLString) {
		t.Error("stringIsValidURL returned true on a string with a non HTTP/HTTPS scheme")
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

func TestHasSchemeAndHost(t *testing.T) {
	if hasSchemeAndHost("/test") {
		t.Error("hasSchemeAndHost returned true when no Scheme or Host were present")
	}
	if hasSchemeAndHost("www.google.com/test") {
		t.Error("hasSchemeAndHost returned true when no Scheme was present")
	}
	if hasSchemeAndHost("http:///test") {
		t.Error("hasSchemeAndHost returned true when no Host was present")
	}
	if !hasSchemeAndHost("http://www.google.com/test") {
		t.Error("hasSchemeAndHost returned false when both Host and Scheme were present")
	}
	if !hasSchemeAndHost("http://google.com/test") {
		t.Error("hasSchemeAndHost returned false when both Host and Scheme were present with no www. prefix")
	}
}

func TestIsSameDomain(t *testing.T) {
	if !isSameDomain("http://www.google.com", "http://www.google.com") {
		t.Error("isSameDomain returned false on matching urls")
	}
	if !isSameDomain("http://www.google.com/test", "http://www.google.com") {
		t.Error("isSameDomain returned false on matching urls with query url having an extension")
	}
	if !isSameDomain("http://www.google.com", "http://www.google.com/test") {
		t.Error("isSameDomain returned false on matching urls with base url having an extension")
	}
	if isSameDomain("https://www.google.com", "http://www.google.com") {
		t.Error("isSameDomain returned true on urls with differnt scheme")
	}
	if isSameDomain("http://www.bing.com", "http://www.google.com") {
		t.Error("isSameDomain returned true on urls with different hosts")
	}
	if isSameDomain("https://www.bing.com", "http://www.google.com") {
		t.Error("isSameDomain returned true on urls with different scheme and hosts")
	}
}

func TestAppendSchemeAndHost(t *testing.T) {
	t.Log(appendSchemeAndHost("http://www.charlie.com", "http://www.charlie.com"))
	t.Log(appendSchemeAndHost("http://www.charlie.com/test", "http://www.charlie.com"))
	t.Log(appendSchemeAndHost("/test", "http://www.charlie.com"))
	t.Log(appendSchemeAndHost("/test.jpg", "http://www.charlie.com"))
}
