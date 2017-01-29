package main

import "log"

func main() {
	url := getInputURL()
	log.Println(url)
	// fmt.Println(page)
}

func getInputURL() string {
	inputURL, err := getInputArguement()
	if err != nil {
		log.Fatalln("input argument error:", err)
	}

	err = stringToValidURL(inputURL)
	if err != nil {
		log.Fatalln("url parsing error:", err)
	}

	return inputURL
}
