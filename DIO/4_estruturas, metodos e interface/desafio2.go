/*
  Parte 2:

	imprimir "PIN" para todos os números que são divisíveis por 3 no range entre 1 e 100
	imprimir "PAN" para todos os números que são divisíveis por 5 no range entre 1 e 100

	imprimir "PINPAN" para todos os números que são divisíveis por 3 e 5 no range entre 1 e 100
*/
package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("PIN+PAN")
			//para saber o número exclua a linha de cima e tire as // dessa:
			//fmt.Println(i, " : PIN+PAN")
		} else if i%3 == 0 {
			fmt.Println(" : PIN")
			//para saber o número:
			//fmt.Println(i, " : PIN")
		} else if i%5 == 0 {
			fmt.Println(i, " : PAN")
			//para saber o número:
			//fmt.Println(i, " : PAN")
		}

	}
}
