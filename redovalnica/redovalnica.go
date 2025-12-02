package redovalnica

import "fmt"

type Student struct {
	Ime     string
	Priimek string
	Ocene   []int
}

func (s Student) String() string {
	return fmt.Sprintf("%s %s", s.Ime, s.Priimek)
}

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

func IzpisVsehOcen(studenti map[string]Student) {
	fmt.Println("REDOVALNICA:")
	for vpisna, student := range studenti {
		fmt.Printf("%s - %s: ", vpisna, student)
		fmt.Println(student.Ocene)
	}
}

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
