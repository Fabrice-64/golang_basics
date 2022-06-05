package main

func main() {
	cards := newDeck()

	h, r := deal(cards, 5)
	r.print()
	h.print()
}
