package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestGetLinks(t *testing.T) {
	t.Log("********* Getting links from ccochrane.com *********")
	resp, _ := http.Get("http://www.ccochrane.com")
	asdf := getLinks(resp)
	for _, a := range asdf {
		t.Log(a)
	}
	t.Log("length: ", len(asdf))

	emptyResponse := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("")),
	}
	links := getLinks(emptyResponse)
	if len(links) != 0 {
		t.Error("getLinks() returned links on an empty reponse body")
	}

	singleLinkResponse := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<a href=\"www.test.com\">test</a>")),
	}
	links = getLinks(singleLinkResponse)
	if len(links) != 1 {
		t.Error("getLinks() returned ", len(links), " links on a reponse body with 1 link")
	} else if links[0] != "www.test.com" {
		t.Error("getLinks() returned an inncorect link on a reponse body with 1 link")
	}

	multipleLinkResponse := &http.Response{
		Body: ioutil.NopCloser(bytes.NewBufferString("<a href=\"www.test1.com\">test1</a><a href=\"www.test2.com\">test2</a>")),
	}
	links = getLinks(multipleLinkResponse)
	if len(links) != 2 {
		t.Error("getLinks() returned ", len(links), " links on a reponse body with 2 links")
	} else if links[0] != "www.test1.com" && links[1] != "www.test2.com" {
		t.Error("getLinks() returned an inncorect link on a reponse body with 2 links")
	}

}

func TestGetAssets(t *testing.T) {
	t.Log("********* Getting assets from ccochrane.com *********")
	resp, _ := http.Get("http://www.ccochrane.com")
	asdf := getAssets(resp)
	for _, a := range asdf {
		t.Log(a)
	}
	t.Log("length: ", len(asdf))
}
