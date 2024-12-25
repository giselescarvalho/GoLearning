package main

import "fmt"

func main() {

	listaa := []float64{100, 99, 98, 97} //lista de notas

	total := 0.0

	for _, valor := range listaa {
		total += valor
	}

	fmt.Println(total / float64(len(listaa))) //imprimir media da sala

}

/* ou tbm:

func media(lista []float64) float64 {
	total := 0.0

	for _, valor := range lista {

		total += valor
	}
	return total / float64(len(lista))
}

func main() {
	lista := []float64{100, 99, 98, 97} //lista de notas
	fmt.Println(media(lista))
}


*/
