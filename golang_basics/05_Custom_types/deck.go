package main

import "fmt"

// Create a new type of deck

type deck []string

func (d deck) print() {
	for idx, card := range d {
		fmt.Println(idx, "-->", card)
	}
}
