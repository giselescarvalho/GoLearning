package main

import (
	"fmt"
	"time"
)

func contador(x int) {
	for i := 0; i < x; i++ {
		fmt.Println(i)
		time.Sleep(time.Second)
	}
}

/*
	func main() {
		contador(10)
		contador(10)
		contador(10)
	}
*/
func main() {
	canal := make(chan string)

	go func() {
		canal <- "opa"
	}()

	msg := <-canal
	fmt.Println(msg) //T1
}

// T1 : linha do func e T2 : primeira linha  do contador -> acessam o mesmo bloco de rotina
// Race Condition -> Condição de corrida
