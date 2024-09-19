package main

import (
	Conta "bank/domain/Conta"
	Dinheiro "bank/domain/Dinheiro"
	"fmt"
)

func main() {
	c := Conta.Conta{
		ID:      "123",
		Titular: "julian",
		Saldo: Dinheiro.Dinheiro{
			Valor: 11.5,
			Moeda: "real"},
	}

	fmt.Println("setembro ", c)
}
