package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("\n-----------------------------------")
	fmt.Println("#1.0 :", retornaumint())
	fmt.Println("#1.1 :", retornauminteumastring())

	//#2
	fmt.Println("\n-----------------------------------")
	si := []int{1, 100, 2, 200, 3, 300, 4, 400, 5, 500, 87}
	fmt.Println("#2 :")
	fmt.Println(somavariadica(si...))
	fmt.Println(somasi(si))

	//#3
	fmt.Println("\n-----------------------------------")
	fmt.Println("#3 :")
	defer fmt.Println(deferdepois(1))
	fmt.Println(deferantes(2))

	//#4 struct pessoa ;
	// criar um método para pessoa que demonstre nome completo e idade ;
	// criar valor de tipo pessoa ;
	// utilizar o método criado para demonstrar esse valor
	fmt.Println("\n-----------------------------------")
	fmt.Println("#4 :")
	aguria := pessoa{
		nome:      "onze",
		sobrenome: "doze",
		idade:     13,
	}
	aguria.demonstrar()

	//#5
	fmt.Println("\n-----------------------------------")
	fmt.Println("#5 :")
	x := quadrado{
		lado: 15,
	}

	y := circulo{
		raio: 16.0,
	}

	medida(x)
	medida(y)
	fmt.Println("https://go.dev/play/p/l4wCteJB2sn")

	//#6 criar função anonima
	fmt.Println("\n-----------------------------------")
	fmt.Println("#6 :")
	func() {
		fmt.Println("Hello, func anonima")
	}()

	//#7 atribuir função a uma variavel
	fmt.Println("\n-----------------------------------")
	fmt.Println("#7 :")
	xpto := func(){
		fmt.Println("Hello, xpto")
	} xpto()

	//#8 função que retorna uma função
	fmt.Println("\n-----------------------------------")
	fmt.Println("#8 :")
	zyq := retornaumafunc()
	zyq()

	//#9 passar uma função como argumento para outra função
	fmt.Println("\n-----------------------------------")
	fmt.Println("#9 :")
	essavaireceberoutrafunc(argumentodefunc)

	
	//#10 closure - função que retorna outra função, onde esta última faz uso de uma variável
	fmt.Println("\n-----------------------------------")
	fmt.Println("#10 :")
	a := ii()
	b := ii()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	//#11 
	fmt.Println("\n-----------------------------------")
	fmt.Println("#11 :")

}


// #1
func retornaumint() int {
	return 10
}

func retornauminteumastring() (int, string) {
	return 20, "string"
}

// #2 recebe um parametro variadico do tipo int e retorna a soma de todos os ints recebidos
func somavariadica(x ...int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}

	return soma
}

// outra funcao que recebe slice of int e retorna
func somasi(x []int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}

	return soma
}

// #3 usando defer
func deferdepois(int) (string, string) {
	s1 := "200"
	s2 := "199"

	return s1, s2
}

func deferantes(int) string {
	s3 := "1"

	return s3
}

// #4
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (FOFAODATV pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é: ", FOFAODATV.nome, FOFAODATV.sobrenome, "\n idade:", FOFAODATV.idade)
	fmt.Println("O nome completo dessa pessoa é: ", FOFAODATV.nome)
}

// #5
type quadrado struct {
	lado int
}
type circulo struct {
	raio     float64
	diametro float64
}

type figura interface{}

type info interface {
	area()
}

func (q quadrado) area() {
	res := q.lado * q.lado
	fmt.Println("area: ", res)
}

func (c circulo) area() {
	resul := math.Pi * 2 * c.raio
	fmt.Println("area: ", resul)
}
func medida(i info) { i.area() }

// #8
func retornaumafunc() func(){
	return func() {
		fmt.Println("Retornando uma func")
	}
}


//#9
func argumentodefunc() {
	fmt.Println("argumento da func")

}

func essavaireceberoutrafunc(x func()) {
	fmt.Println("Prestenção")
	x()
}

//#10 closure
func ii() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}