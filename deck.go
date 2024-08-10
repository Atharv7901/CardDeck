package main

import (
	"fmt"
	"os"
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

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0666)
}

func newDeckfromFile(fileName string) deck {
	bs, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)

}