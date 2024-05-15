package main

import "fmt"

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Println("Worker %d received %d \n", workerID, x)
	}
}apackage main

import (
	"fmt"
	"time"
)

func worker(workerID int, data chan int) {
	for x := range data {
		fmt.Println("Worker %d received %d \n", workerID, x)
		time.Sleep(time.Second)
	}
}

/*
func main() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {
		canal <- i
	}
}    >>>>> deadlock!
*/

func main() { //T1
	canal := make(chan int)

	qtdWorkers := 50

	for i := 0; i < qtdWorkers; i++ {
		go worker(i, canal)
	}

	for i := 0; i < 100; i++ {
		canal <- i
	}
}

/* com 1>>>>>>  1 0
; 1 1 ;  1 2 ;  1 3 ;  1 4  ; 1 5  ;  1 6 ;  1 7 ;  1 8 ;  1 9
*/

/* com 2>>>>>> Worker %d received %d
2 0 ;  1 1  ;   1 2   ;   2 3  ;  2 4  ;  1 5  ;  2 6 ;  1 7  ;  1 8 */


func main() {
	canal := make(chan int)
	for i := 0; i < 10; i++ {
		canal <- i
	}
}
