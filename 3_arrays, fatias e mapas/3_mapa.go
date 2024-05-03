package main

import "fmt"

func main() {

	x := make(map[string]int)
	x["chave"] = 10
	fmt.Println(x["chave"])

	y := make(map[int]int)
	y[1] = 20
	y[2] = 30
	fmt.Println(y[1], y[2])

	elemento := make(map[string]string)
	elemento["H"] = "hidrogênio"
	elemento["He"] = "hélio"
	elemento["Li"] = "lítio"
	fmt.Println(elemento["Li"], elemento["He"])

}
