package main

import (
	"errors"
	"flag"
)

func getInputArguement() (string, error) {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		err := errors.New("Please specify a single starting http or https web address as first arguement")
		return "", err
	}

	return args[0], nil
}
