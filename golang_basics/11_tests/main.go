package main

import "log"

func main() {
	cards := newDeck()
	cards.print()
	_, hand := cards.deal(7)
	hand.print()
	r := cards.saveToFile("my_cards")
	if r != nil {
		log.Fatalln(r)
	}
	s := readFile("my_cards")
	s.shuffle()
	s.print()
}
