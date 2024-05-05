package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
} //campos relacionados à pessoa

func main() {
	fmt.Println(pessoa{"Ana", 54})
	fmt.Println(pessoa{"Joao", 55})
}
