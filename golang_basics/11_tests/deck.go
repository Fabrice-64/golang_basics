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

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func readFile(filename string) deck {
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
