package main

import "fmt"

type retangulo struct {
	comprimento, altura int
}

// area
func (r retangulo) area() int {
	return r.comprimento * r.altura
}

// perimetro
func (r retangulo) perimetro() int {
	return 2*r.comprimento + 2*r.altura
}

// main
func main() {
	r := retangulo{comprimento: 10, altura: 5}

	fmt.Println("Área: ", r.area())
	fmt.Println("Perimetro:", r.perimetro())

}
