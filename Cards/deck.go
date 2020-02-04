package main

import "fmt"

// type "deck"
type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}


func newDeck() deck {
	cards := deck{}
	cardSuites := []string{"Spades"."Diamonds","Hearts","Clubs"}
	cardValues := []string{"Ace","Two","Three","Four"}

	for i, suit := range cardSuites {
		for j, value := range cardValues{
			cards = append(cards, suit+" of "+"value")
		}
	}

	return cards

}
