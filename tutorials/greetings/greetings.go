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

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	// Loop through received slice of names
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
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
