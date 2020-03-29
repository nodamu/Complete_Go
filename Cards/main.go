package main

func main() {
	// cards := newDeckFromFile("hello.txt")
	// fmt.Println(cards.toString())
	// cards.saveToFile("hello.txt")
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
