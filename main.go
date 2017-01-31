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
	url, err := getInputURL()
	if err != nil {
		log.Fatalln(err)
	}

	// crawl from the starting URL
	err = crawl(url)
	if err != nil {
		log.Fatalln(err)
	}
}
