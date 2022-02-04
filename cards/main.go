package main

func main() {
	// var card string = "Full house"
	// card := "Full House"
	// card := "ace of spades" -> ini akan error
	// card = "ace of spades"
	// card := newCard()
	// fmt.Println(card)
	// cards := deck{"Ace of Diamonds", kartuBaru()}
	// cards = append(cards, "Straight")

	cards := kartuBaru()

	cards.tampilkan()
}
