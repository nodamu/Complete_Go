package main

import "fmt"

type person struct {
	firstname string
	lastname  string
}

func main() {
	nick := person{firstname: "Nicholas", lastname: "Adamu"}
	fmt.Print(nick)
}
