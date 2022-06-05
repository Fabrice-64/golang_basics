package main

import "fmt"

type deck []string

func newDeck() {
	cards := deck{}
	cardSuits := []string{"Hearts", "Diamonds", "Clubs", "Spades"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	fmt.Println(cards)

}

func (d deck) print() {
	for idx, card := range d {
		fmt.Println(idx, "-->", card)
	}
}
