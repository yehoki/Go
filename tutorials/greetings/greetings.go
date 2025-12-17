package greetings

import (
	"errors"
	"fmt"
)

// Hello returns a greeting for the person provided
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hi, %v. Welcome - here's a greeting!", name)
	return message, nil
}
