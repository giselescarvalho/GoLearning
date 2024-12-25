package main

import (
	"fmt"
	"sort"
)

type Dados struct {
	Nome  string
	Idade int
}

// interface
type ParaNome []Dados

func (ps ParaNome) Len() int {
	return len(ps)
}

// Less: item na posição i é < que item na posição j
func (ps ParaNome) Less(i, j int) bool {
	return ps[i].Nome < ps[j].Nome
}

// swap troca os itens
func (ps ParaNome) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

func main() {
	criancas := []Dados{
		{"Julia", 5},
		{"Joao", 8},
	}
	sort.Sort(ParaNome(criancas))
	fmt.Println(criancas)
}
