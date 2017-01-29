package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"
)

func init() {
	httpClient = &httpTestClient{}
}

func TestGetPage(t *testing.T) {
	_, err := getPageBytes("returnError")
	if err == nil {
		t.Error("GetPage() returned no error on a http client Get request error")
	}

	respBytes, err := getPageBytes("returnValid")
	if err != nil {
		t.Error("GetPage() returned an error on a http client Get request than returned correctly")
	}
	if string(respBytes) != "hi" {
		t.Error("GetPage() returned incorrect bytes from body of a valid responce")
	}

}

type httpTestClient struct {
}

func (*httpTestClient) Get(url string) (*http.Response, error) {
	if url == "returnError" {
		return nil, errors.New("client Error")
	}

	if url == "returnValid" {
		resp := &http.Response{
			Body: ioutil.NopCloser(bytes.NewBufferString("hi")),
		}
		return resp, nil
	}

	return nil, nil
}
