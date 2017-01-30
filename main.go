package main

import "log"

func main() {
	url := getInputURL()
	log.Println("Starting at URL: ", url)
	//recursiveGet

}

func getInputURL() string {
	inputURL, err := getInputArguement()
	if err != nil {
		log.Fatalln("input argument error:", err)
	}

	if !stringIsValidURL(inputURL) {
		log.Fatalln("input URL is not valid, please enter a http/https URL")
	}

	return inputURL
}
