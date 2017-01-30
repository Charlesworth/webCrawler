package main

import (
	"io"
	"log"
	"os"
)

var jsonPrinter io.Writer

func main() {
	jsonPrinter = os.Stdout
	url := getInputURL()
	crawl(url)
}

func getInputURL() string {
	inputURL := getInputArguement()

	if !stringIsValidURL(inputURL) {
		log.Fatalln("input URL is not valid, please enter a http/https URL")
	}
	return inputURL
}
