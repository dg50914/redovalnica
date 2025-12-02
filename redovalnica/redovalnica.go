// Paket redovalnica ponuja preprost način upravljanja z ocenami študentov, prikazovanjem redovalnice in izračunavanjem povprečij.
//
// Primer uporabe:
//
//	studenti := map[string]redovalnica.Student{
//		"123": {
//			Ime:     "Franc",
//			Priimek: "Horvat",
//			Ocene:   []int{10, 9, 10, 9, 10, 10, 9},
//		},
//		"456": {
//			Ime:     "Ana",
//			Priimek: "Kovačič",
//			Ocene:   []int{6, 7, 6, 7, 6, 7, 6, 7},
//		},
//		"789": {
//			Ime:     "Janez",
//			Priimek: "Novak",
//			Ocene:   []int{4, 5, 6, 5, 4, 5, 4, 5, 4},
//		},
//		"234": {
//			Ime:     "Zdravo",
//			Priimek: "Neocenjen",
//			Ocene:   []int{7, 8, 9},
//		},
//	}
//
//	redovalnica.IzpisVsehOcen(studenti)
//
//	redovalnica.DodajOceno(studenti, "123", 8, minOcena, maxOcena)
//	redovalnica.IzpisVsehOcen(studenti)
//
//	redovalnica.IzpisiKoncniUspeh(studenti, stOcen)
package redovalnica

import "fmt"

// Student predstavlja enega študenta
type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func (s Student) String() string {
	return fmt.Sprintf("%s %s", s.Ime, s.Priimek)
}

// DodajOceno danemu študentu doda oceno, če se ta nahaja med min in max
func DodajOceno(studenti map[string]Student, vpisnaStevilka string, ocena int, min, max int) {
	if !(ocena >= min && ocena <= max) {
		fmt.Printf("Ocena %d ni v intervalu [%d, %d]\n", ocena, min, max)
		return
	}

	student, ok := studenti[vpisnaStevilka]

	if !ok {
		fmt.Printf("Študent z vpisno številko %s ne obstaja\n", vpisnaStevilka)
		return
	}

	student.Ocene = append(student.Ocene, ocena)
	studenti[vpisnaStevilka] = student
	fmt.Printf("Ocena %d dodana študentu/ki %s\n", ocena, student)
}

// IzpisVsehOcen izpiše vse študente z vpisnimi številkami in njihove ocene
func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s: ", vpisna, student)
		fmt.Println(student.Ocene)
	}
}

// povprecje izračuna povprečno oceno študenta.
// Če ima študent manj kot minStOcen ocen, vrne 0.
// Če študenta ne more najti, vrne -1.
func povprecje(studenti map[string]Student, vpisnaStevilka string, minStOcen int) float64 {
	student, ok := studenti[vpisnaStevilka]
	if !ok {
		return -1
	}

	if len(student.Ocene) < minStOcen {
		return 0
	}

	avg := 0.0
	for _, ocena := range student.Ocene {
		avg += float64(ocena)
	}

	avg /= float64(len(student.Ocene))
	return avg
}

// IzpisiKoncniUspeh izpiše vse študente, njihove povprečene ocene in dosežen uspeh
// Če je povprečje < 6, je uspeh neuspešen
// Če je povprečje >= 6 in povprečje < 9, je uspeh povprečen
// Če je povprečje >=, je uspeh odličen
func IzpisiKoncniUspeh(studenti map[string]Student, minStOcen int) {
	fmt.Println("KONČNI USPEH:")
	for vpisna, student := range studenti {
		studentAvg := povprecje(studenti, vpisna, minStOcen)

		fmt.Printf("%s: povprečna ocena %.1f -> ", student, studentAvg)

		switch {
		case studentAvg >= 9:
			fmt.Printf("Odličen študent!\n")
		case studentAvg < 9 && studentAvg >= 6:
			fmt.Printf("Povprečen študent\n")
		case studentAvg < 6:
			fmt.Printf("Neuspešen študent\n")
		}
	}
}
