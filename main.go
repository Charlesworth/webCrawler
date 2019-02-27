package main

import (
	"log"
	"os"

	"github.com/charlesworth/webCrawler/jsonPrinter"
)

func main() {
	if len(os.Args) < 2 || !validURL(os.Args[1]) {
		log.Fatal("Please specify a valid http or https URL as first arguement")
	}
	url := os.Args[1]

	// set the jsonPrinter to output to std out
	jsonPrinter.Writter = os.Stdout

	// crawl from the starting URL
	err := crawl(url)
	if err != nil {
		log.Fatalln(err)
	}
}
