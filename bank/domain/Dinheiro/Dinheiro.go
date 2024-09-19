package domain

type Dinheiro struct {
	Valor float64
	Moeda string
}

func (d Dinheiro) Somar(outro Dinheiro) Dinheiro {
	return Dinheiro{
		Valor: d.Valor + outro.Valor,
		Moeda: d.Moeda,
	}
}

func (d Dinheiro) Subtrair(outro Dinheiro) Dinheiro {
	return Dinheiro{
		Valor: d.Valor - outro.Valor,
		Moeda: d.Moeda,
	}
}
