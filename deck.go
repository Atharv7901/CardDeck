package main

import "fmt"

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