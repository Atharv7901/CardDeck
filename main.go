package main

import "fmt"

func main() {
	fmt.Println("Welcome to Deck of Cards!")
	cards := newDeck()

	hand, remainingHand := deal(cards, 5)

	cards.print()
	fmt.Println("-------------------------------")
	hand.print()
	remainingHand.print()

	cards.saveToFile("my_file")
}