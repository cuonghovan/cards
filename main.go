package main

func main() {
	cards := deck{"Ba ro", newCard()}
	cards = append(cards, "Chin bich")
	cards.print()
}

func newCard() string {
	return "At co"
}
