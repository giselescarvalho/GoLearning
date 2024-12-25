package main

import "fmt"

func main() {

	//#1 de 1 a 10k
	for x := 1; x <= 10000; x++ {
		fmt.Printf("%v _ ", x)

	}
	fmt.Println("-----------------------------------")

	//#2
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v _ ", i)
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n ", i)
		}

	}

	fmt.Println("-----------------------------------")

	//#3
	fmt.Println("#3 :")
	anoqueeunasci := 1996
	anoquerocontar := 2077

	for anoqueeunasci <= anoquerocontar {
		fmt.Printf("%v ; ", anoqueeunasci)
		anoqueeunasci++
	}
	fmt.Println("-----------------------------------")

	//#4
	fmt.Println("#4 :")

	for {
		if anoqueeunasci == anoquerocontar {
			break
		}
		fmt.Printf("%v ; ", anoqueeunasci)
		anoqueeunasci++
	}
	fmt.Println("-----------------------------------")

	//#5
	fmt.Println("#5 :")
	for div := 10; div <= 100; div++ {
		fmt.Printf("%v/4 tem resto: %v ; \n", div, div%4)
	}
	fmt.Println("\n-----------------------------------")

	//#6 demonstrar funcionamento de if
	fmt.Println("#6 :")
	temmacanamacieira := true
	if temmacanamacieira {
		fmt.Println("Se tem maçã na macieira\nÉ só subir pra pegar")
	}
	fmt.Println("\n-----------------------------------")

	//#7
	fmt.Println("#7 :")
	quantasmacastem := 10
	if quantasmacastem == 100 {
		fmt.Println("Se tem maçã na macieira\nÉ só subir pra pegar")
	} else if quantasmacastem < 100 {
		fmt.Println("Corre que tá acabando")
	} else {
		fmt.Println("Demorou")
	}
	fmt.Println("\n-----------------------------------")

	//#8
	fmt.Println("#8 :")
	switch {
	case quantasmacastem == 100:
		fmt.Println("Se tem maçã na macieira\nÉ só subir pra pegar!")
	case quantasmacastem == 50:
		fmt.Println("Corre que tá acabando...")
	case quantasmacastem == 10:
		fmt.Println("Acabando...")
	case quantasmacastem == 0:
		fmt.Println("Demorou")
	default:
		fmt.Println("Pode ser que ainda tenha...")
	}

	fmt.Println("\n-----------------------------------")

	//#9
	fmt.Println("#9 :")
	esporteFavorito := "Tênis"
	switch esporteFavorito {
	case "Futebol":
		fmt.Println("Welcome to Brasil")
	case "Tênis":
		fmt.Println("Hoobie caro...")
	case "Rugby":
		fmt.Println("é...")
	case "Corrida":
		fmt.Println("meus pesâmes")
	default:
		fmt.Println("Pode ser que ainda tenha salvação...")
	}
}
