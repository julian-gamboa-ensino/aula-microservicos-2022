package service

import (
	"projeto-posto/controller"
	"projeto-posto/model"
)

type PostoService interface {
	ProcessarVeiculos(
		veiculos []model.Veiculo,
		modelos map[string]model.Modelo) ([]model.Saida, []model.Saida)
}

type postoService struct{}

func NewPostoService() PostoService {
	return &postoService{}
}

func (ps *postoService) ProcessarVeiculos(veiculos []model.Veiculo, modelos map[string]model.Modelo) ([]model.Saida, []model.Saida) {
	var saidasEtanol, saidasGasolina []model.Saida
	tempoFilaEtanol, tempoFilaGasolina := 0.0, 0.0

	for _, veiculo := range veiculos {
		modelo := modelos[veiculo.Modelo]
		var combustivel string

		if modelo.ConsumoEtanol > 0 && modelo.CustoPorKmEtanol < modelo.CustoPorKmGasolina {
			combustivel = "ETANOL"
			tempoFilaEtanol += modelo.TempoCarregamentoEtanol
			saida := model.Saida{
				TempoCarregamento: tempoFilaEtanol,
				Menssagem:         controller.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoEtanol),
			}
			saidasEtanol = append(saidasEtanol, saida)
		} else {
			combustivel = "GASOLINA"
			tempoFilaGasolina += modelo.TempoCarregamentoGasolina
			saida := model.Saida{
				TempoCarregamento: tempoFilaGasolina,
				Menssagem:         controller.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoGasolina),
			}
			saidasGasolina = append(saidasGasolina, saida)
		}
	}

	return saidasEtanol, saidasGasolina
}
