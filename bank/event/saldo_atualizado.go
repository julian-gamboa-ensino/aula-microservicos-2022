package event

import (
	Dinheiro "bank/domain/Dinheiro"
)

type SaldoAtualizado struct {
	ContaID string
	Saldo   Dinheiro.Dinheiro
}
