package main

import (
	"fmt"
	"golang-basics/golang_modules/greetings"
)

func main() {
	greet := greetings.Hello("Gladys")
	fmt.Println(greet)
}
