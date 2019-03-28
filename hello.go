package main

import (
	"fmt"
	"flag"
)

func main() {
	var (
		help bool
	)
	
	flag.BoolVar(&help, "help", false, "Print this help message.")
	flag.Parse()
	
	if !help {
		fmt.Println("Hello, world!")
	}
}
