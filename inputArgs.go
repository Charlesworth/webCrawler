package main

import (
	"flag"
	"fmt"
	"os"
)

func getStartingPage() string {
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		fmt.Println("No arguement supplied")
		fmt.Println("Please specify a single starting http/https web address as first arguement")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("More than 1 arguement supplied")
		fmt.Println("Please specify a single starting http/https web address as first arguement")
		os.Exit(1)
	}

	return args[0]
}
