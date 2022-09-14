package main

import "fmt"

// Go Playground: https://play.golang.org/
func main() {
	// var card string = "Ace of Spades"
	card := "Ace of Spades" // Only use := when declaring new variable
	card = "Five of Diamonds"

	card = newCard()

	fmt.Println(card)
}

func newCard() string {
	return "Eight of Diamonds"
}
