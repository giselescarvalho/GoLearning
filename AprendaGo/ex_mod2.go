package main

import (
	"fmt"
)

var a int
var b string
var c bool

type hotdog int

var h hotdog

var iota int

func main() {
	//#1
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("-----------------------------------")

	//#2

	fmt.Println("#2 :\nint: '", a, "'", "\nstring: '", b, "'", "\nbool: '", c, "'")

	fmt.Println("-----------------------------------")
	//#3
	s := fmt.Sprintf("#3:\na: %v\nb: %v\nc: %v", a, b, c)
	fmt.Println(s)

	fmt.Println("-----------------------------------")
	//#4
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	h := 42
	fmt.Println(h)

	fmt.Println("-----------------------------------")
	//#5
	iota = int(h)
	fmt.Println(iota)
	fmt.Printf("%T\n", iota)

}
