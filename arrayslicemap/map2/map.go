package main

import "fmt"

func main() {
	funcsESalarios := map[string]float64{ // tipo de map e o valor
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Junior":   1200.0,
	}

	funcsESalarios["Rafael Filho"] = 1350.0
	delete(funcsESalarios, "Inexistenta")

	for nome, salario := range funcsESalarios {
		fmt.Println(nome, salario)
	}
}
