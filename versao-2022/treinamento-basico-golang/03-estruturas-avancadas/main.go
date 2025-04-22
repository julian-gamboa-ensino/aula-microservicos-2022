package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade int
}

func (p Pessoa) Saudacao() {
	fmt.Printf("Olá, meu nome é %s e eu tenho %d anos.\n", p.Nome, p.Idade)
}

func main() {
	pessoa := Pessoa{"Julian", 35}
	pessoa.Saudacao()
}
