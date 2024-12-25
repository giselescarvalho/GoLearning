package main

import "fmt"

type Pessoa struct {
	nome             string
	sobrenome        string
	saboresfavoritos []string
}

//3
type caminhoneta struct {
	Veiculo
	tracaoNasQuatro bool
}

type sedan struct {
	Veiculo
	modeloLuxo bool
}

type Veiculo struct {
	portas int
	cor    string
}

func main() {

	//#1

	pessoa1 := Pessoa{
		nome:             "Anitta",
		sobrenome:        "Larissa",
		saboresfavoritos: []string{"Chocolate Branco", "Flocos", "Choco Amargo"},
	}

	pessoa2 := Pessoa{"Lorem", "Ipsum", []string{"Sabão em pó", "De Feijão", "Açaí"}}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)

	fmt.Println(pessoa1.nome, pessoa1.sobrenome, ":")
	for _, v := range pessoa1.saboresfavoritos {
		fmt.Println("\t", v)
	}

	//#2 com map
	fmt.Println("\n-----------------------------------")
	fmt.Println("#2 :")
	apa := make(map[string]Pessoa)
	mapa["Ipsum"] = Pessoa{"Joao", "Ipsum", []string{"morango", "choco"}}

	fmt.Println("----------------")

	for _, valor := range mapa {
		fmt.Println("O sabor favorito do ", valor.nome, "é", valor.saboresfavoritos)
	}

	fmt.Println("-- melhorando --")

	for _, valor := range mapa {
		fmt.Println("O sabor favorito do ", valor.nome, "é")
		for _, valor := range valor.saboresfavoritos {
			fmt.Println("-", valor)
		}
	}

	//#3
	fmt.Println("\n-----------------------------------")
	fmt.Println("#3 :")
	veiculoCaminhoneta := caminhoneta{Veiculo{4, "transparente"}, true}
	sedanHomem := sedan{Veiculo{4, "cinza"}, false}

	fmt.Println(veiculoCaminhoneta)
	fmt.Println(veiculoCaminhoneta.tracaoNasQuatro)
	fmt.Println(sedanHomem)
}

	//#4
	fmt.Println("\n-----------------------------------")
	fmt.Println("#4 :")

	any := struct {
		esporte                    string
		atualcampeao               []string
		paisesGanhadoresOlimpiadas map[string]string
	}{
		esporte:      "Judô",
		atualcampeao: []string{"Japão", "Diadema"},
		paisesGanhadoresOlimpiadas: map[string]string{
			"Japão":  "8 vezes",
			"Brasil": "1 vez",
		},
	}

	fmt.Println(any)


}
