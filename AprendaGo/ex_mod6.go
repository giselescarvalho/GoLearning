package main

import "fmt"

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
	fmt.Println(exercicio5(x))

	//#6
	fmt.Println("\n-----------------------------------")
	fmt.Println("#6 :")

	//#7
	fmt.Println("\n-----------------------------------")
	fmt.Println("#7 :")

	//#8
	fmt.Println("\n-----------------------------------")
	fmt.Println("#8 :")

	//#9
	fmt.Println("\n-----------------------------------")
	fmt.Println("#9 :")
}

//#1
func retornaumint() int {
	return 10
}

func retornauminteumastring() (int, string) {
	return 20, "string"
}

//#2 recebe um parametro variadico do tipo int e retorna a soma de todos os ints recebidos
func somavariadica(x ...int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}

	return soma
}

//outra funcao que recebe slice of int e retorna
func somasi(x []int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}

	return soma
}

//#3 usando defer
func deferdepois(int) (string, string) {
	s1 := "200"
	s2 := "199"

	return s1, s2
}

func deferantes(int) string {
	s3 := "1"

	return s3
}

//#4
type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

func (FOFAODATV pessoa) demonstrar() {
	fmt.Println("O nome completo dessa pessoa é: ", FOFAODATV.nome, FOFAODATV.sobrenome, "\n idade:", FOFAODATV.idade)
	fmt.Println("O nome completo dessa pessoa é: ", FOFAODATV.nome)
}

//#5
type qudrado struct {
	lado int
}
type circulo struct {
	raio     int
	diametro int
}

func (q quadrado) area() {
	res := q.lado
	fmt.Println("area: ", res)
}

func exercicio5() {

}

type figura interface{}
