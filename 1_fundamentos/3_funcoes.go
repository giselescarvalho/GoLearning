package main

import (
	"fmt"
)

func hello(nome string) {
	fmt.Printf("Hello ", nome, "!")
}

func sum(a, b int) int {
	//return ("A soma entre ", a, " e ", b, " é: ", a + b)
	return a + b
}

/*func main() {
	hello("Gisele")
}*/
