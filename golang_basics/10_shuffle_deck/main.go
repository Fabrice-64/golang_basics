package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	s := newDeckFromFile("my_cards")
	fmt.Println(s)
}
