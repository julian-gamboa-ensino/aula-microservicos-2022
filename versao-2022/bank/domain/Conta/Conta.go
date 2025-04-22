package domain

import (
	Dinheiro "bank/domain/Dinheiro"
	"errors"
)

type Conta struct {
	ID      string
	Titular string
	Saldo   Dinheiro.Dinheiro
}

func (c *Conta) RealizarDeposito(valor Dinheiro.Dinheiro) error {
	if valor.Valor <= 0 {
		return errors.New("valor do depósito deve ser maior que zero")
	}
	c.Saldo = c.Saldo.Somar(valor)
	return nil
}

func (c *Conta) RealizarSaque(valor Dinheiro.Dinheiro) error {
	if valor.Valor > c.Saldo.Valor {
		return errors.New("saldo insuficiente")
	}
	c.Saldo = c.Saldo.Subtrair(valor)
	return nil
}

func (c *Conta) RealizarTransferencia(destino *Conta, valor Dinheiro.Dinheiro) error {
	if err := c.RealizarSaque(valor); err != nil {
		return err
	}
	return destino.RealizarDeposito(valor)
}
