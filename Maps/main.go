package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "ff35f",
		"blue":  "hhhtl",
		"green": "4bjhg",
	}

	printMap(colors)
	// fmt.Print(a)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
