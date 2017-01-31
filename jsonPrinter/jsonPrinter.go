package jsonPrinter

import (
	"encoding/json"
	"fmt"
	"io"
)

// Writter is used to pass jsonPrinter a writter to target its output
var Writter io.Writer

// firstEntry is used to determine if its the first page to be printed, prefixing the array bracket
var firstEntry = true

// Page is the struct used to marshal page URL and asset info to JSON
type Page struct {
	URL    string   `json:"url"`
	Assets []string `json:"assets"`
}

// End is called to close the JSON array when all entries have been printed
func End() {
	fmt.Fprintln(Writter, "]")
}

// PrintPage is used print a page URL and assets to std out in JSON format
func PrintPage(url string, assets []string) {

	// first marshal the URL and assets into JSON format
	page := Page{
		url,
		assets,
	}
	res, err := json.Marshal(page)
	// return and print nothing if a marshalling error occured
	if err != nil {
		return
	}

	// if its the first entry, prefix the opening square bracket, else append a comma to the previos entity
	if firstEntry {
		fmt.Fprint(Writter, "[")
		firstEntry = false
	} else {
		fmt.Fprintln(Writter, ",")
	}

	// print the page in its JSON format
	fmt.Fprint(Writter, string(res))
}

// Reset is used to reset the printer for a new array to be printed
func Reset() {
	firstEntry = true
}
