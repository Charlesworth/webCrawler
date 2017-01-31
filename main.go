package main

import (
	"log"
	"os"
	"scrapper/jsonPrinter"
)

func main() {
	// set the jsonPrinter to output to std out
	jsonPrinter.Writter = os.Stdout

	// get the starting URL from input arguement
	url := getInputURL()

	// crawl from the starting URL
	err := crawl(url)
	if err != nil {
		log.Fatalln(err)
	}
}

func getInputURL() string {
	inputURL := getInputArguement()

	if !stringIsValidURL(inputURL) {
		log.Fatalln("input URL is not valid, please enter a http/https URL")
	}
	return inputURL
}
