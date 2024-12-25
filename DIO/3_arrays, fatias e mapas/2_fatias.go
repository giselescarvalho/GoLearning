package main

import "fmt"

/*func main() {
	fatia := make([]float64, 4)
	fmt.Println(fatia)
		//[0 0 0 0]
}*/

//fatia
/*func main() {
	an := [4]float64{1, 5, 25, 250}
	fatia := an[1:4]
	fmt.Println(fatia)
		//[5 25 250]
}*/

//acrescentar com append
/*func main() {
	fatia1 := []int{1, 5, 10, 25}
	fatia2 := append(fatia1, 4, 5)
	fmt.Println(fatia1, fatia2)
		//[1 5 10 25] [1 5 10 25 4 5]
}*/

//copiar
func main() {
	fatia1 := []int{1, 5, 10, 25}
	fatia2 := make([]int, 2) //copiar 2 elementos apenas
	//fatia2 := make([]int, 6) //acrescenta zeros
	copy(fatia2, fatia1)
	fmt.Println(fatia1, fatia2)
}
