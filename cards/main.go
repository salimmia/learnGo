package main

import "fmt"

func main() {
	cards := newDeck()

	// cards := newDeck()

	// cards.saveToFile("my_cards")

	fs, se, th := deal(cards, 10)
	fs.print()
	fmt.Println()
	se.print()
	fmt.Println()
	th.print()
	fmt.Println()

	cards = newDeckFromFile("my_cards")

	cards.print()

	cards.shuffle()

	cards.print()
}