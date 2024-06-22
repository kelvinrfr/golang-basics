package greetings

import (
	"errors"
	"fmt"
)

// HIello returns a greeting for the named person
func Hello(name string) (string, error) {
	// if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("message should not be empty")
	}

	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
