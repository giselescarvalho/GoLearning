package main

import "fmt"

/*var x float64
var y int
func main() {
	x = 10.8
	y = 10
	fmt.Printf("X : %g %T, e Y: %g %T", x,x, y,y)
}*/

var a int = 20
var b string = "Nome"

func main() {
	var c float64 = float64(a)
	fmt.Printf("o valor de C é:  %g  e o valor de B é:  %s .", c, b)
}
