package main

import (
	"fmt"

	"github.com/dg50914/redovalnica/redovalnica"
)

func main() {

	studenti := map[string]redovalnica.Student{
		"123": {
			Ime:     "Franc",
			Priimek: "Horvat",
			Ocene:   []int{10, 9, 10, 9, 10, 10, 9},
		},
		"456": {
			Ime:     "Ana",
			Priimek: "Kovačič",
			Ocene:   []int{6, 7, 6, 7, 6, 7, 6, 7},
		},
		"789": {
			Ime:     "Janez",
			Priimek: "Novak",
			Ocene:   []int{4, 5, 6, 5, 4, 5, 4, 5, 4},
		},
		"234": {
			Ime:     "Zdravo",
			Priimek: "Neocenjen",
			Ocene:   []int{7, 8, 9},
		},
	}

	fmt.Println("Začetno stanje:")
	redovalnica.IzpisRedovalnice(studenti)
	fmt.Println()

	fmt.Println("Dodajanje ocen:")
	fmt.Println()
	redovalnica.DodajOceno(studenti, "123", 8)
	redovalnica.IzpisRedovalnice(studenti)
	fmt.Println()

	redovalnica.DodajOceno(studenti, "123", 15)
	redovalnica.IzpisRedovalnice(studenti)
	fmt.Println()

	redovalnica.DodajOceno(studenti, "abc", 4)
	redovalnica.IzpisRedovalnice(studenti)
	fmt.Println()

	fmt.Println()
	redovalnica.IzpisiKoncniUspeh(studenti)
}
