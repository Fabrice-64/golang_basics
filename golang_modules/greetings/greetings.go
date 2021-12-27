package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello to you %v", name)
	return message
}
