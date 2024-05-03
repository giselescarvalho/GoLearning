package main

import "fmt"

const ebuilicaoK = 373.0

func main() {

	var tempK = ebuilicaoK
	var tempC = (tempK - 273)

	fmt.Println("A temperatura da água em °F é: ", tempK)
	fmt.Println("A temperatura da água em °C é: ", tempC)

}
