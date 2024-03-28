package main

func main() {
	cards := newDeckFromFile("my_cards")
	// cards.saveToFile("my_cards")

	cards.print()

}
