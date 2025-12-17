package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

// Hello returns a greeting for the person provided
func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// randomFormat returns a random message out of a set of messages.
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Great to see you %v",
		"Hail %v, hello",
		"Yo %v",
		"Heyo matey %v",
	}

	return formats[rand.Intn(len(formats))]
}
