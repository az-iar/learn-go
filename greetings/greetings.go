package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// if no name given, return error with message
	if name == "" {
		return "", errors.New("empty name")
	}

	// if name received, return a value
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

func init() {
	rand.NewSource(time.Now().UnixNano())
}

func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
