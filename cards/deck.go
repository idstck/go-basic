package main

import "fmt"

// membuat tipe data deck
// di mana sebagai slice of strings
type deck []string

func kartuBaru() deck {
	cards := deck{}

	jenisKartu := []string{"Sekop", "Wajik", "Hati", "Keriting"}
	nilaiKartu := []string{"As", "Dua", "Tiga", "Empat"}

	for _, jenis := range jenisKartu {
		for _, nilai := range nilaiKartu {
			cards = append(cards, nilai+" "+jenis)
		}
	}

	return cards
}

func (d deck) tampilkan() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
