package main

func main() {
	cards := newDeck()
	cards.print()
	_, hand := cards.deal(7)
	hand.print()
}
