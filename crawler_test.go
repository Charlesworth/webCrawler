package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"scrapper/jsonPrinter"
)

func TestCrawl(t *testing.T) {
	buffer := setJSONPrinterForTest()

	t.Log("== test crawl with a unreachable starting address")
	crawl("http://")
	output := buffer.String()
	if output != "" {
		t.Error("test crawl with a non-existant starting address should return no JSON")
	}

	t.Log("== test crawl with a response with no links or assets")
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}))
	buffer = setJSONPrinterForTest()
	crawl(testServer.URL)
	testServer.Close()
	output = buffer.String()
	if output != "[{\"url\":\""+testServer.URL+"\",\"assets\":[]}]\n" {
		t.Error("test crawl with a page starting address with no links or assets should return A single entity with no assets")
		t.Error("returned:", output)
	}

	t.Log("== test crawl with a response with a single internal domain link and asset")
	testServer = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "<img href=\"/test.jpg\">test1</img><a href=\"/test\">test</a>")
	}))
	buffer = setJSONPrinterForTest()
	crawl(testServer.URL)
	testServer.Close()
	output = buffer.String()
	if output != "[{\"url\":\""+testServer.URL+"\",\"assets\":[\""+testServer.URL+"/test.jpg\"]},\n{\"url\":\""+testServer.URL+"/test\",\"assets\":[\""+testServer.URL+"/test.jpg\"]}]\n" {
		t.Error("test crawl with a page starting address with a single link and sinlge asset should return two entitys with one asset each")
		t.Error("returned:", output)
	}
}

func setJSONPrinterForTest() *bytes.Buffer {
	buffer := new(bytes.Buffer)
	jsonPrinter.Writter = buffer
	jsonPrinter.Reset()
	return buffer
}
