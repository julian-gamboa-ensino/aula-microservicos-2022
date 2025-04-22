package service

import (
	Conta "bank/domain/Conta"
	Dinheiro "bank/domain/Dinheiro"
	repository "bank/repository"
)

type CriarContaService struct {
	RepositorioDeContas repository.RepositorioDeContas
}

func (s *CriarContaService) CriarConta(titular string, saldoInicial Dinheiro.Dinheiro) (*Conta.Conta, error) {
	conta := &Conta.Conta{
		ID:      "gerarIDUnico()",
		Titular: titular,
		Saldo:   saldoInicial,
	}

	if err := s.RepositorioDeContas.Salvar(conta); err != nil {
		return nil, err
	}
	return conta, nil
}
