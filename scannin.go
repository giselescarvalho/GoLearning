package main

import "fmt"

func main() {
    var input string
    fmt.Println("Digite algo: ")
    // Lê uma linha de entrada do usuário e armazena na variável 'input'
    fmt.Scanln(&input)
    fmt.Println("Você digitou: ", input)
}