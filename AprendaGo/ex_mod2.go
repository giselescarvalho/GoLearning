package main

import (
	"fmt"
)

// #2
const i int = 10
const j = 10

// #6
const (
	_ = 2024 + iota //ignorando ano atual
	anoquevem
	anodepois
	anodepsdeps
	anodoisvintoito
)

func main() {
	//#1

	x := 100
	//decimal
	//hexadecimal
	//binário

	fmt.Println("#1 :")
	fmt.Printf("%d, %#x, %b\n", x, x, x)
	fmt.Println("-----------------------------------")

	//#2
	fmt.Println("#2 :")
	a := (10 == 10) //true
	b := (10 != 10) //false
	c := (10 <= 10) //true
	d := (10 < 10)  //false
	e := (10 >= 10) //true
	f := (10 > 10)  //false

	fmt.Printf(" %v,\n %v,\n %v,\n %v,\n %v,\n %v,\n", a, b, c, d, e, f)
	fmt.Println("-----------------------------------")

	//#3
	fmt.Println("#3 :")
	fmt.Println("tipado: ", i, "\nnão tipado: ", j)
	fmt.Println("-----------------------------------")

	//#4 deslocar
	bb := 300
	fmt.Printf("%d\t%b\t%#x\n", bb, bb, bb)
	bc := bb << 1
	fmt.Printf("%d\t%b\t%#x\n", bc, bc, bc)
	fmt.Println("-----------------------------------")

	//#5 raw string literal
	k := `imprindo
		quase 
			um
			json
		.`
	fmt.Println("#5 :")
	fmt.Printf(k)
	fmt.Println("\n-----------------------------------")

	//#6 raw string literal
	fmt.Println("#6 : ")
	fmt.Println(anoquevem,
		anodepois,
		anodepsdeps,
		anodoisvintoito)

}
