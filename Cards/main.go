package main

func main() {
	cards := deck{"Ace of spade", newCard(), newCard()}
	cards = append(cards, "Jack")

	cards.print()

}

func newCard() deck {
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
