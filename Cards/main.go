package main

func main() {
	cards := deck{"Ace of spade", newCard(), newCard()}
	cards = append(cards, "Jack")

	cards.print()

}

func newCard() string {
	return "Diamond"
}
