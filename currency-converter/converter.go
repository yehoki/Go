package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to the currency converter")
	for {
		var input string
		fmt.Scanln(&input)
		if input == "c" {
			break
		}
		fmt.Print(input)
	}
}
