package main

import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds","Clubs"}
	cardValue := []string {"Ace", "One", "Two", "Three"}

	for _, suit := range cardSuits{
		for _, value := range cardValue {
			cards = append(cards, value + " of "+ suit)
		}
	}

	return cards
}

func (d deck) print(){
	for i, value := range d{
		fmt.Println(i,value)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}