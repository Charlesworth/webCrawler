package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func getPageBytes(u *url.URL) ([]byte, error) {

	resp, err := http.Get(u.String())
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	resp.Body.Close()
	return bytes, nil
}
