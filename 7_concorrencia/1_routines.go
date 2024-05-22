// funções paralelas/concorrentes
package main

import "fmt"

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, " : ", i)
	}
}

func main() {
	go f(0 + 1)
	var escreva string
	fmt.Scanln(&escreva)
}
