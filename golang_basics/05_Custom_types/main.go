package main

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "2 of Spikes")
	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}

// To run: go run main.go deck.go
