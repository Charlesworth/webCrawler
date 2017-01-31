package main

import (
	"errors"
	"flag"
)

func getInputArguement() (string, error) {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		return "", errors.New("Error: Please specify a single starting http or https web address as first arguement")
	}

	return args[0], nil
}

func getInputURL() (string, error) {
	inputURL, err := getInputArguement()
	if err != nil {
		return "", err
	}

	if !stringIsValidURL(inputURL) {
		return "", errors.New("input URL is not valid, please enter a http/https URL")
	}
	return inputURL, nil
}
