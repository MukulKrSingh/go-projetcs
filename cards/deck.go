package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Creating our own type
type deck []string

func newDeck() deck {

	cards := deck{}

	cardSuits := []string{"Heart", "Spade", "Diamond", "Clove"}
	cardValues := []string{"Ace", "One", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards

}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Receiver function somewhat similar to extension in dart
// !This is not parameter being passed to function
func (d deck) print() {
	for i, cards := range d {
		fmt.Println(i, cards)
	}
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(fileName string) error {
	return os.WriteFile(fileName, []byte(d.toString()), 0644)
}

func newDeckFromFile(fileName string) deck {
	bs, error := os.ReadFile(fileName)

	if error != nil {
		fmt.Println("Error: ", error)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		randPos := r.Intn(len(d) - 1)

		d[i], d[randPos] = d[randPos], d[i]
	}
}
