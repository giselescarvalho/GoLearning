package main

import (
	"log"
	"net/http"
)

func main() {
	fileSer := http.FileServer(http.Dir("./"))
	http.Handle("/", fileSer)

	log.Print("Listening on: 3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal()
	}
}

//http://localhost:3000/aula4.html
