package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards.txt")
	fmt.Println(cards)
}
