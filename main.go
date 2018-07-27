package main

func main() {
	// cards := newDeck()
	// cards.writeToFile("cards")
	cards1 := newDeckFromFile("cards")
	cards1.print()
}
