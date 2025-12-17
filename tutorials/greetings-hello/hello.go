package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Get a greeting msg and print it.
	// Example here, running the following commands:
	// 1. go mod edit -replace example.com/greetings=../greetings
	// This is to ensure we can locate a local dependency
	// 2. go mod tidy
	// Ensuring the module is registered correctly
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
