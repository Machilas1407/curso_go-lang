package main

import "fmt"

func main() {
	i := 1

	var p *int = nil // equivale ao null

	p = &i // pegando o endereço da variável e apontando para o ponteiro
	*p++   // desreferenciando (pegando o valor)

	// Go não tem aritimética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}
