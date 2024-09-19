package domain

import (
	Dinheiro "bank/domain/Dinheiro"
	"testing"
)

func TestRealizarDeposito(t *testing.T) {
	conta := Conta{
		ID:      "1",
		Titular: "Teste",
		Saldo:   Dinheiro.Dinheiro{Valor: 100, Moeda: "real"},
	}

	valorDeposito := Dinheiro.Dinheiro{Valor: 50, Moeda: "real"}
	err := conta.RealizarDeposito(valorDeposito)

	if err != nil {
		t.Errorf("Erro inesperado ao realizar depósito: %v", err)
	}

	if conta.Saldo.Valor != 150 {
		t.Errorf("Saldo após depósito incorreto. Esperado: 150, Obtido: %f", conta.Saldo.Valor)
	}

	// Teste de depósito com valor inválido (zero ou negativo)
	valorDepositoInvalido := Dinheiro.Dinheiro{Valor: -10, Moeda: "real"}
	err = conta.RealizarDeposito(valorDepositoInvalido)

	if err == nil {
		t.Errorf("Esperado erro ao realizar depósito com valor inválido")
	}
}

func TestRealizarSaque(t *testing.T) {
	conta := Conta{
		ID:      "2",
		Titular: "Teste",
		Saldo:   Dinheiro.Dinheiro{Valor: 100, Moeda: "real"},
	}

	valorSaque := Dinheiro.Dinheiro{Valor: 30, Moeda: "real"}
	err := conta.RealizarSaque(valorSaque)

	if err != nil {
		t.Errorf("Erro inesperado ao realizar saque: %v", err)
	}

	if conta.Saldo.Valor != 70 {
		t.Errorf("Saldo após saque incorreto. Esperado: 70, Obtido: %f", conta.Saldo.Valor)
	}

	// Teste de saque com saldo insuficiente
	valorSaqueInsuficiente := Dinheiro.Dinheiro{Valor: 150, Moeda: "real"}
	err = conta.RealizarSaque(valorSaqueInsuficiente)

	if err == nil {
		t.Errorf("Esperado erro ao realizar saque com saldo insuficiente")
	}
}

func TestRealizarTransferencia(t *testing.T) {
	contaOrigem := Conta{
		ID:      "3",
		Titular: "Teste Origem",
		Saldo:   Dinheiro.Dinheiro{Valor: 100, Moeda: "real"},
	}

	contaDestino := Conta{
		ID:      "4",
		Titular: "Teste Destino",
		Saldo:   Dinheiro.Dinheiro{Valor: 50, Moeda: "real"},
	}

	valorTransferencia := Dinheiro.Dinheiro{Valor: 40, Moeda: "real"}
	err := contaOrigem.RealizarTransferencia(&contaDestino, valorTransferencia)

	if err != nil {
		t.Errorf("Erro inesperado ao realizar transferência: %v", err)
	}

	if contaOrigem.Saldo.Valor != 60 {
		t.Errorf("Saldo da conta origem após transferência incorreto. Esperado: 60, Obtido: %f", contaOrigem.Saldo.Valor)
	}

	if contaDestino.Saldo.Valor != 90 {
		t.Errorf("Saldo da conta destino após transferência incorreto. Esperado: 90, Obtido: %f", contaDestino.Saldo.Valor)
	}

	// Teste de transferência com saldo insuficiente
	valorTransferenciaInsuficiente := Dinheiro.Dinheiro{Valor: 200, Moeda: "real"}
	err = contaOrigem.RealizarTransferencia(&contaDestino, valorTransferenciaInsuficiente)

	if err == nil {
		t.Errorf("Esperado erro ao realizar transferência com saldo insuficiente")
	}
}
