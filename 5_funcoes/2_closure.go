package main

import "fmt"

func main() {

	x := 10

	numero := func() int {
		x++
		return x
	}

	fmt.Println(numero()) //11
	fmt.Println(numero()) //12
}

//função chama variável que está em outra função
