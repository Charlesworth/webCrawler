package main

import "log"

func main() {
	url := getInputURL()
	log.Println(url)
}

func getInputURL() string {
	inputURL, err := getInputArguement()
	if err != nil {
		log.Fatalln("input argument error:", err)
	}

	err = stringIsValidURL(inputURL)
	if err != nil {
		log.Fatalln("url parsing error:", err)
	}

	return inputURL
}
