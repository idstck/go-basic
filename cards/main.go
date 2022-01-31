package main

import "fmt"

func main() {
	// var card string = "Full house"
	// card := "Full House"
	// card := "ace of spades" -> ini akan error
	// card = "ace of spades"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Lima Hati"
}
