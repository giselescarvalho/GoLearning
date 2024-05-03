package main

import "fmt"

func main() {

	x := 3
	y := 3
	if x == 2 || x == 3 && y%2 == 0 {
		fmt.Println("X pode ser 2 ou 3")
	} else {
		fmt.Println("X não é 2 ou 3 OU Y não é par")
	}
}
