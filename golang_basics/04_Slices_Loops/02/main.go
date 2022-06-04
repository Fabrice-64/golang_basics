package main

import "fmt"

func main() {
	cards := []string{"2 of Spades", newCard()}
	fmt.Println(cards)
	cards = append(cards, "4 of Diamonds")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Ace of Spades"
}
