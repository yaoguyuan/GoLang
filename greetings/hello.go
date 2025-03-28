package greetings

import "fmt"

func SayHello(name string) string {
	message := fmt.Sprintf("Hello, %v. Welcome!", name)
	return message
}
