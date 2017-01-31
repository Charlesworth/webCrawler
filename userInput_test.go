package main

import (
	"os"
	"testing"
)

func TestGetInputURL(t *testing.T) {
	// no args supplied
	os.Args = []string{"scrapper"}
	_, err := getInputURL()
	if err == nil {
		t.Error("getInputURL did not return an error when no url arguement was supplied")
	}

	// more than 1 arg supplied
	os.Args = []string{"scrapper", "http://www.google.com", "http://www.google.com"}
	_, err = getInputURL()
	if err == nil {
		t.Error("getInputURL did not return an error when more than 1 url arguement was supplied")
	}

	// URL supplied is not valid
	os.Args = []string{"scrapper", "google.com"}
	_, err = getInputURL()
	if err == nil {
		t.Error("getInputURL did not return an error when more than 1 url arguement was supplied")
	}

	// valid single URL supplied
	os.Args = []string{"scrapper", "http://www.test.com"}
	testval, err := getInputURL()
	if err != nil {
		t.Error("getInputURL returned an error on a single valid URL arguement")
	}
	if testval != "http://www.test.com" {
		t.Error("getInputURL returned an inncorrect URL than that which was supplied")
	}
}
