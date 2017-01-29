package main

import "log"

func main() {
	inputString, err := getInputArguement()
	if err != nil {
		log.Fatalln("input argument error:", err)
	}

	url, err := stringToValidURL(inputString)
	if err != nil {
		log.Fatalln("url parsing error:", err)
	}

	log.Println(url)
	// fmt.Println(page)
}
