package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Hearts", "Diamonds", "Spades", "Clubs"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four", "Five", "Six", "Seven"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for idx, card := range d {
		fmt.Println("Card:", idx, " --> ", card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[handSize:], d[:handSize]
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return strings.Split(string(bs), ", ")
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		j := r.Intn(len(d) - 1)
		d[i], d[j] = d[j], d[i]
	}
}
