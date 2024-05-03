package main

import "fmt"

//temperatura da água em graus Kelvin
const ebuilicaoK = 373.0

func main() {

	tempK := ebuilicaoK
	tempC := (tempK - 273)

	fmt.Println("A temperatura da água em °F é: ", tempK)
	fmt.Println("A temperatura da água em °C é: ", tempC)

}
