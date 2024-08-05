package main

import "fmt"

func main() {
	fmt.Println("Welcome to Deck of Cards!")
	cards := newDeck()

	cards.print()
}