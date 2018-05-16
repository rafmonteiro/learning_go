package main

import "fmt"

// create a new type of `deck` which is
// a slace of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds"}
	cardValues := []string{"Ace", "Two"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

// Print a deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
