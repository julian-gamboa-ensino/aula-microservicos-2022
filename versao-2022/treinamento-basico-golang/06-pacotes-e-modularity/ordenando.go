package main

import "fmt"

func main() {
    idades := map[string]int{
        "Alice": 25,
        "Bob": 30,
        "Charlie": 20,
    }

    maiorValor := 0
    chaveMaiorValor := ""
    for nome, idade := range idades {
        if idade > maiorValor {
            maiorValor = idade
            chaveMaiorValor = nome
        }
    }

    fmt.Printf("A pessoa com a maior idade é %s com %d anos\n", chaveMaiorValor, maiorValor)
}