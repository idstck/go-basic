package main

import "fmt"

// membuat tipe data deck
// di mana sebagai slice of strings
type deck []string

func (d deck) tampilkan() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
