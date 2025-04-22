package domain

import (
	"testing"
)

func TestSomar(t *testing.T) {
	dinheiro1 := Dinheiro{Valor: 50, Moeda: "real"}
	dinheiro2 := Dinheiro{Valor: 30, Moeda: "real"}

	resultado := dinheiro1.Somar(dinheiro2)

	if resultado.Valor != 80 || resultado.Moeda != "real" {
		t.Errorf("Soma incorreta. Esperado: 80 real, Obtido: %f %s", resultado.Valor, resultado.Moeda)
	}
}

func TestSubtrair(t *testing.T) {
	dinheiro1 := Dinheiro{Valor: 50, Moeda: "real"}
	dinheiro2 := Dinheiro{Valor: 30, Moeda: "real"}

	resultado := dinheiro1.Subtrair(dinheiro2)

	if resultado.Valor != 20 || resultado.Moeda != "real" {
		t.Errorf("Subtração incorreta. Esperado: 20 real, Obtido: %f %s", resultado.Valor, resultado.Moeda)
	}
}
