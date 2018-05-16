package main

// create a new type of `deck` which is
// a slace of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds"}
	cardValues := []string{"Ace", "Two"}

	for i, suit := range cardSuits {
		for j, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
