package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{ // string chave e o map é o valor
		"G": {
			"Gabriela Silva": 15456.78,
			"Guga Pereira":   8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letras, funcs := range funcsPorLetra {
		fmt.Println(letras, funcs)
	}
}
