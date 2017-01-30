package main

import (
	"flag"
	"log"
)

func getInputArguement() string {
	flag.Parse()
	args := flag.Args()

	if len(args) != 1 {
		log.Fatalln("Error: Please specify a single starting http or https web address as first arguement")
	}

	return args[0]
}
