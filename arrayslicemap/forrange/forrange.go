package main

import "fmt"

func main() {
	numeros := [...]int{1, 2, 3, 4, 5} // compilador conta!

	for i, numeros := range numeros {
		fmt.Printf("%d) %d\n", i+1, numeros)
	}

	for _, num := range numeros { /* o _ serve para ignorar o indice, se não tiver
		ele vai apresentar o indice e não o valor */
		fmt.Println(num)
	}
}
