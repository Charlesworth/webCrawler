package main

import (
	"io/ioutil"
	"net/http"
)

var httpClient client

type client interface {
	Get(string) (*http.Response, error)
}

func init() {
	httpClient = &http.Client{}
}

func getPageBytes(url string) ([]byte, error) {

	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	return bytes, err
}
