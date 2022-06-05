package main

func main() {
	cards := deck{"Ace of Spades", newCard()}
	cards = append(cards, "3 of Hearts")
	cards.print()
	newDeck()
}

func newCard() string {
	return "3 of Spades"
}

// To run:  go run main.go deck.go
