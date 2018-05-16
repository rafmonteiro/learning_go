package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainingDeck := deal(cards, 2)
	hand.print()
	fmt.Println("The remaininig deck is: ")
	remainingDeck.print()
}
