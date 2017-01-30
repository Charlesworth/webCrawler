package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawl(t *testing.T) {
	buffer := new(bytes.Buffer)
	jsonPrinter = buffer

	t.Log("== test crawl with a non-existant starting address")
	crawl("http://")
	output := buffer.String()
	if output != "[\n]\n" {
		t.Error("test crawl with a non-existant starting address did not return an empty JSON array '[]'")
	}

	t.Log("== test crawl with a response with no links or assets")
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	buffer = new(bytes.Buffer)
	jsonPrinter = buffer
	crawl(testServer.URL)
	output = buffer.String()
	if output != "[\n{\"url\":\""+testServer.URL+"\",\"assets\":[]},\n]\n" {
		t.Error("WRONG")
	}
	testServer.Close()

	t.Log("== test crawl with a response with a single internal domain link and asset")
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<img href=\"/test.jpg\">test1</img><a href=\"/test\">test</a>")
	}))
	buffer = new(bytes.Buffer)
	jsonPrinter = buffer
	crawl(testServer.URL)
	output = buffer.String()
	if output != "[\n{\"url\":\""+testServer.URL+"\",\"assets\":[\""+testServer.URL+"/test.jpg\"]},\n{\"url\":\""+testServer.URL+"/test\",\"assets\":[\""+testServer.URL+"/test.jpg\"]},\n]\n" {
		t.Error("WRONG")
	}
	testServer.Close()
}

func TestPrintJSON(t *testing.T) {
	buffer := new(bytes.Buffer)
	jsonPrinter = buffer

	// test printJSON with a valid URL and Asset slice
	printJSON("hello", []string{"goodbye"})
	output := buffer.String()
	if output != "{\"url\":\"hello\",\"assets\":[\"goodbye\"]}" {
		t.Error("printJSON is output is incorrect")
	}
}
