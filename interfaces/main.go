package main

import (
	"fmt"
)

//thats right, no fields.
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sp := spanishBot{}

	printGreeting(eb)
	printGreeting(sp)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
func (englishBot) getGreeting() string {
	// some crazy logic for english greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	// more crazy loginc for spanish greeting
	return "Hola!"
}
