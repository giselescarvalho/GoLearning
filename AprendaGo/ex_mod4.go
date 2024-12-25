package main

import "fmt"

func main() {

	//#1
	fmt.Println("#1 :")
	array := [5]int{25, 50, 100, 200, 400}
	for _, v := range array {
		fmt.Println(v, ";")
	}
	//tipo do array
	fmt.Printf("%T", array)

	//#2 literal composta
	fmt.Println("\n-----------------------------------")
	fmt.Println("#2 :")
	slice := []int{4, 8, 12, 16, 20, 24, 28, 32, 36, 40}

	for indice, s := range slice {
		fmt.Println("para o indice ", indice, ":", s, ";")
	}

	fmt.Printf("%T", slice)

	//#3 slicing para demonstrar valores
	fmt.Println("\n-----------------------------------")
	fmt.Println("#3 :")
	slice2 := []int{4, 8, 12, 16, 20, 24, 28, 32, 36, 40}

	fmt.Println("do primeiro ao terceiro", slice2[0:3])
	fmt.Println("do quinto ao último", slice2[4:])
	fmt.Println("do segundo ao sétimo, incluindo sétimo", slice2[1:7])
	fmt.Println("do terceiro ao penúltimo", slice2[2:9])

	fmt.Println("\t\t\t", slice2[2:len(slice2)])

	//#4
	fmt.Println("\n-----------------------------------")
	fmt.Println("#4 :")
	slicex := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("inicializando\t\t\t", slicex)

	slicex = append(slicex, 52)
	fmt.Println("depois de add o ultimo elemento:", slicex)

	slicex = append(slicex, 53, 54, 55)
	fmt.Println("add uma unica declaração:\t", slicex)

	umapartedasliceouy := []int{56, 57, 58}
	slicex = append(slicex, umapartedasliceouy...)
	fmt.Println("depois de add o ultimo elemento:", slicex)

	//#5
	fmt.Println("\n-----------------------------------")
	fmt.Println("#5 :")
	slicez := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("inicializando\t\t\t", slicez)

	slicez = append(slicez[:3], slicez[len(slicez)-4:]...)
	fmt.Println("vendo primeiro teste\t\t", slicez)

	slicez = append(slicez[:3], slicez[3:]...)
	fmt.Println("testando a msm coisa \t\t", slicez)

	//#6 demonstrar len e cap ; sem usar range
	fmt.Println("\n-----------------------------------")
	fmt.Println("#6 :")
	estados := make([]string, 26, 27)
	fmt.Println("antes de add: ")
	fmt.Println(len(estados))
	fmt.Println(cap(estados))

	estados[0], estados[1], estados[2], estados[3], estados[4], estados[5], estados[6], estados[7], estados[8], estados[9], estados[10], estados[11],
		estados[12], estados[13], estados[14], estados[15], estados[16], estados[17], estados[18], estados[19], estados[20],
		estados[21], estados[22], estados[23], estados[24], estados[25] = "Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
		"Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco",
		"Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins"
	//

	fmt.Println("depois de add: ") //nao roda com len 25 pq dá bug na alocação pois nao estamos usando append
	fmt.Println(len(estados))
	fmt.Println(cap(estados))

	states := make([]string, 25, 28)
	states = []string{"Acre", "Alagoas", "Amapá", "Amazonas", "Bahia", "Ceará", "Espírito Santo", "Goiás",
		"Maranhão", "Mato Grosso", "Mato Grosso do Sul", "Minas Gerais", "Pará", "Paraíba", "Paraná", "Pernambuco",
		"Piauí", "Rio de Janeiro", "Rio Grande do Norte", "Rio Grande do Sul", "Rondônia", "Roraima", "Santa Catarina",
		"São Paulo", "Sergipe", "Tocantins"}
	fmt.Println(" ---- outra forma ----  ") //aqui dá pra alocar mais do que a capacidade de 25
	fmt.Println(len(states))
	fmt.Println(cap(states))

	//#7 multidimensional
	fmt.Println("\n-----------------------------------")
	fmt.Println("#7 :")
	nomesobrenomehobbie := [][]string{
		[]string{"", "", ""}, // poderia ser []string{} mas dá erro no range
		[]string{"Chis", "Burguer", "Comer Hamburguer"},
		[]string{"Jana", "Winchester", "Jogar Dominó"},
	}

	fmt.Println(nomesobrenomehobbie)

	for _, nome := range nomesobrenomehobbie {
		fmt.Println(nome[0])
		for _, nomes := range nome {
			fmt.Println("\t", nomes)
		}
	}

	//#8
	fmt.Println("\n-----------------------------------")
	fmt.Println("#8 :")
	mapa := make(map[string]string)

	mapa = map[string]string{
		"Silva_Maria": "Jogar CyberPunk 2077",
		"Doe_John":    "Jogar Fute",
	}

	mape := map[string][]string{
		"pereira_joana": []string{"jogar nintendo", "jogar among us", "jogar fortnite"},
		"joao_jose":     []string{"comer pizza", "fazer tatuagem", "palavras cruzadas"},
	}

	fmt.Println("_", mapa)
	fmt.Println("_", mape)

	for indicenumero, valor := range mape {
		fmt.Println("\nIndice: ", indicenumero, "gosta de", valor[0], "e", valor[1])
	}

	fmt.Println("\n---\nmais apresentavel\n---")
	for k, v := range mape {
		fmt.Println(k)
		for kk, vv := range v {
			fmt.Println("\t", kk, "-", vv)
		}
	}

	//#9
	fmt.Println("\n-----------------------------------")
	fmt.Println("#9 :")
	//mape := map[string][]string{
	//	"pereira_joana": []string{"jogar nintendo", "jogar among us", "jogar fortnite"},
	//	"joao_jose":     []string{"comer pizza", "fazer tatuagem", "palavras cruzadas"},
	//}
	mape["anitta_larissa"] = []string{"cantar", "aprender novos idiomas"}

	fmt.Println("---\n\n---")
	for k, v := range mape {
		fmt.Println(k)
		for kk, vv := range v {
			fmt.Println("\t", kk, "-", vv)
		}
	}

	fmt.Println("\n-----------------------------------")
	fmt.Println("#10 :")
	delete(mape, "joao_jose")

	fmt.Println("---------\ndepois de eliminar joao\n---------")
	for k, v := range mape {
		fmt.Println(k)
		for kk, vv := range v {
			fmt.Println("\t", kk, "-", vv)
		}
	}

}
