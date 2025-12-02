package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/dg50914/redovalnica/redovalnica"
	"github.com/urfave/cli/v3"
)

func main() {

	cmd := &cli.Command{
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:  "stOcen",
				Value: 6,
				Usage: "Min število potrebovanih ocen",
			},
			&cli.IntFlag{
				Name:  "minOcena",
				Value: 0,
				Usage: "Najmanjša možna ocena",
			},
			&cli.IntFlag{
				Name:  "maxOcena",
				Value: 10,
				Usage: "Največja možna ocena",
			},
		},
		Action: func(ctx context.Context, cmd *cli.Command) error {
			stOcen := cmd.Int("stOcen")
			minOcena := cmd.Int("minOcena")
			maxOcena := cmd.Int("maxOcena")

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
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()

			fmt.Println("Dodajanje ocen:")
			fmt.Println()
			redovalnica.DodajOceno(studenti, "123", 8, minOcena, maxOcena)
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()

			redovalnica.DodajOceno(studenti, "123", 15, minOcena, maxOcena)
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()

			redovalnica.DodajOceno(studenti, "abc", 4, minOcena, maxOcena)
			redovalnica.IzpisVsehOcen(studenti)
			fmt.Println()

			fmt.Println()
			redovalnica.IzpisiKoncniUspeh(studenti, stOcen)

			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
