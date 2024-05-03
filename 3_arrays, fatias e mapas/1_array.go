package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 5.3
	x[1] = 8
	x[2] = 9.0

	var total float64 = 0
	for i := 0; i < 5; i++ {
		fmt.Println("iteração número ", i, " x: ", x)
		total += x[i]
	}

	fmt.Println(" é ", total/5)
}
