package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "King", "Queen", "Jack"}

	for _, i := range cardSuits {
		for _, j := range cardValues {
			cards = append(cards, j+" of "+i)
		}
	}

	return cards
}

func (d deck) print() {
	for idx, card := range d {
		fmt.Println("Card:", idx, "-->", card)
	}
}

func (d deck) deal(handsize int) (deck, deck) {
	return d[handsize:], d[:handsize]
}

func (d deck) toString() string {
	return strings.Join(d, ", ")
}
