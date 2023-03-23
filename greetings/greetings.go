package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// if no name given, return error with message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name received, return a value
	message := fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
