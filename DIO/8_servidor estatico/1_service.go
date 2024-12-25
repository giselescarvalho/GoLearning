package main

import (
	"fmt"
	"net/http"
)

func ola(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "ola\nola\nola\nola\nola\nola\nola\nola\n")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintln(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	//rota da funcao
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalho", cabecalhos)

	//http://localhost:8090/cabecalho
	http.ListenAndServe(":8090", nil)
}
