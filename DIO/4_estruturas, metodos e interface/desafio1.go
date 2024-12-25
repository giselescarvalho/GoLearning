/* 
  Parte 1:
	imprimir todos os números que são divisíveis por 3 no range entre 1 e 100
*/
package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, " é divisível por 3\n")
		}
	}

}
