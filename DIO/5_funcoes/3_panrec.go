package main

import "fmt"

func main() {

	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Panico")

}

//panic - erro do programador/ timeout
//recover - interrompe o panic
