// funções paralelas/concorrentes
package main

import (
	"fmt"
	"time"
)

func pingar(c chan string) { //palavra reservada para Canal: chan
	// sem limite por isso ; ;  e vai enviar receber do canal com <-
	for i := 0; ; i++ {
		c <- "para sair tecle Enter\n\nping\n"
	}
}

func imprimir(c chan string) {
	for {
		msg := <-c
		fmt.Print(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {

	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var escreva string
	fmt.Scanln(&escreva)
}
