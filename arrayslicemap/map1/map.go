package main

import (
	"fmt"
)

func main() {
	// var aprovados map[int]string // A chave vai ser o int e o valor uma string
	// mapas devem ser inicializados
	aprovados := make(map[int]string) // tipo de map e o valor

	aprovados[12345678978] = "Maria"
	aprovados[12465131231] = "Pedro"
	aprovados[58765646546] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[58765646546])
	delete(aprovados, 58765646546)
	fmt.Println(aprovados[58765646546])

}
