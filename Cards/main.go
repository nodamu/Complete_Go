package main

import "fmt"

func main() {
	card := deck{"Ace of spade", newCard(), newCard()}
	card = append(card, "Jack")

	for index, card := range card {
		fmt.Println(index, card)
	}

}

func newCard() string {
	return "Diamond"
}
