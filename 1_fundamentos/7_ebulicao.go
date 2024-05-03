package main

import "fmt"

const ebuilicaoF float64 = 212.0

func main() {

	var tempF float64 = ebuilicaoF
	var tempC float64 = (tempF - 32) * 5 / 9

	fmt.Println("A temperatura da água em °F é: ", tempF)
	fmt.Println("A temperatura da água em °C é: ", tempC)

}
