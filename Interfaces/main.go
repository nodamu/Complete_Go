package main

import "fmt"

// bot interface
// every type that implements getGreeting
// with a string return type is a "member" of bot
type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	sb := spanishBot{}
	eb := englishBot{}

	printGreeting(sb)
	printGreeting(eb)

}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// Pretending to write custom logic for englishBot

	return "Hello there"
}

func (eb spanishBot) getGreeting() string {
	return "Hola"
}
