package greetings

import (
	"errors"
	"fmt"
)

func SayHello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message, nil
}
