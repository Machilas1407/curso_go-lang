package main

import "fmt"

func main() {
	fmt.Printf("Mesma")
	fmt.Printf(" linha.")

	fmt.Println("Nova")
	fmt.Println("linha.")

	x := 3.141516

	//fmt.Println("O valor de x é " + x)  isso não funciona em Go

	xs := fmt.Sprint(x)
	fmt.Println("o valor de x é " + xs)
	fmt.Println("o valoe de x é", x, "!!!")
	fmt.Printf("o valor de x é %.2f", x) /* o %f informa que ele vai receber um valor do tipo float
	e o 2 a quantidade de casas decimais que ele vai apresentar, printf quer dizer formatado */

	a := 1
	b := 1.9999
	c := false
	d := "opa"

	fmt.Printf("\n %d %f %.1f %t %s", a, b, b, c, d) // d é numero inteiro, f float, t bolean, s string
	fmt.Printf("\n %v %v %v %v", a, b, c, d)         // o v é uma forma mais generica, imprime todos os valores
}
