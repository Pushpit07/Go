package main

import "fmt"

// Go Playground: https://play.golang.org/
func main() {
	cards := newDeck()
	cards.shuffle()

	cards.saveToFile("saved_cards")
	hand, remainingCards := deal(cards, 3)

	fmt.Println("Hand dealt:")
	hand.print()
	fmt.Println("\nRemaining cards:")
	remainingCards.print()
}
