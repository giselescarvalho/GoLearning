//Para consultar: http://pokeapi.co/api/v2/pokedex/kanto/

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Response struct {
	name    string    `json: "name`
	Pokemon []Pokemon `json: "pokemon_entries"`
}

type Pokemon struct {
	numero  int            `json: "entry_number`
	especie PokemonSpecies `json: "pokemon_species"`
}

type PokemonSpecies struct {
	name string `json: name`
}

func main() {

	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/") //mapeia o Http

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	//lendo
	responseData, err := ioutil.ReadAll(response.Body) //lendo o corpo
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(responseData))

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.name)
	fmt.Println((responseObject.Pokemon))

	for _, pokemon := range responseObject.Pokemon {
		fmt.Println(pokemon.especie.name)
	}

}
