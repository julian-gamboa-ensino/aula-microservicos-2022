package infra

import (
	Conta "bank/domain/Conta"
	"errors"
)

type RepositorioDeContasMemoria struct {
	contas map[string]*Conta.Conta
}

func NovoRepositorioDeContasMemoria() *RepositorioDeContasMemoria {
	return &RepositorioDeContasMemoria{
		contas: make(map[string]*Conta.Conta),
	}
}

func (r *RepositorioDeContasMemoria) Salvar(conta *Conta.Conta) error {
	r.contas[conta.ID] = conta
	return nil
}

func (r *RepositorioDeContasMemoria) BuscarPorID(contaID string) (*Conta.Conta, error) {
	conta, existe := r.contas[contaID]
	if !existe {
		return nil, errors.New("conta não encontrada")
	}
	return conta, nil
}
