package greetings

import "fmt"

// Hello returns a greeting for the person provided
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome - here's a greeting!", name)
	return message
}
