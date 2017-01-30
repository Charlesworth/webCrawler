package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestDuffTest(t *testing.T) {
	t.Log("********* Getting links from ccochrane.com *********")
	resp, _ := http.Get("http://www.ccochrane.com")
	asdf := getLinks(resp)
	for _, a := range asdf {
		t.Log(a)
	}
	t.Log("length: ", len(asdf))

	t.Log("******** Getting assets from ccochrane.com ********")
	resp, _ = http.Get("http://www.ccochrane.com")
	asdf = getAssets(resp)
	for _, a := range asdf {
		t.Log(a)
	}
	t.Log("length: ", len(asdf))
}

func TestGetLinks(t *testing.T) {
	testRequest, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		t.Error("unable to make test request")
	}

	emptyResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("")),
	}
	links := getLinks(emptyResponse)
	if len(links) != 0 {
		t.Error("getLinks() returned links on an empty reponse body")
	}

	singleLinkResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("<a href=\"www.test.com\">test</a>")),
	}
	links = getLinks(singleLinkResponse)
	if len(links) != 1 {
		t.Error("getLinks() returned ", len(links), " links on a reponse body with 1 link")
	} else if links[0] != "www.test.com" {
		t.Error("getLinks() returned an inncorect link on a reponse body with 1 link")
	}

	multipleLinkResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("<a href=\"www.test1.com\">test1</a><a href=\"www.test2.com\">test2</a>")),
	}
	links = getLinks(multipleLinkResponse)
	if len(links) != 2 {
		t.Error("getLinks() returned ", len(links), " links on a reponse body with 2 links")
	} else if links[0] != "www.test1.com" && links[1] != "www.test2.com" {
		t.Error("getLinks() returned an inncorect link on a reponse body with 2 links")
	}

}

func TestGetAssets(t *testing.T) {
	testRequest, err := http.NewRequest("GET", "http://www.google.com", nil)
	if err != nil {
		t.Error("unable to make test request")
	}

	emptyResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("")),
	}
	assets := getAssets(emptyResponse)
	if len(assets) != 0 {
		t.Error("getAssets() returned links on an empty reponse body")
	}

	singleAbsoluteURLLinkResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("<img href=\"http://www.google.com/test.jpg\">test</img>")),
	}
	assets = getAssets(singleAbsoluteURLLinkResponse)
	if len(assets) != 1 {
		t.Error("getAssets() returned ", len(assets), " links on a reponse body with 1 link")
	} else if assets[0] != "http://www.google.com/test.jpg" {
		t.Error("getAssets() returned an inncorect link on a reponse body with 1 link")
	}

	singleLinkNoSchemeAndHostResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("<img href=\"/test.jpg\">test1</img>")),
	}
	assets = getAssets(singleLinkNoSchemeAndHostResponse)
	if len(assets) != 1 {
		t.Error("getAssets() returned ", len(assets), " links on a reponse body with 1 links")
	} else if assets[0] != "http://www.google.com/test.jpg" {
		t.Error("getAssets() returned an inncorect link on a link with no Scheme or Host")
	}

	multipleLinkResponse := &http.Response{
		Request: testRequest,
		Body:    ioutil.NopCloser(bytes.NewBufferString("<img href=\"/test.jpg\">test1</img><script src=\"http://www.google.com/test.js\"></script>")),
	}
	assets = getAssets(multipleLinkResponse)
	if len(assets) != 2 {
		t.Error("getAssets() returned ", len(assets), " links on a reponse body with 2 links")
	} else if assets[0] != "http://www.google.com/test.jpg" && assets[1] != "http://www.google.com/test.js" {
		t.Error("getAssets() returned an inncorect link on a link with no Scheme or Host")
	}
}
