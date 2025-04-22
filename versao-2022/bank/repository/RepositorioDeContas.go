package repository

import (
	Conta "bank/domain/Conta"
)

type RepositorioDeContas interface {
	Salvar(conta *Conta.Conta) error
	BuscarPorID(contaID string) (*Conta.Conta, error)
}
