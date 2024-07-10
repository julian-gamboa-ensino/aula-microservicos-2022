package main

import (
	"fmt"
	"os"
	"projeto-posto/controllers"
	"projeto-posto/datareader"
	"projeto-posto/models"
)

func main() {
	csvReader := &datareader.CSVReader{
		VeiculosFile: "data/veiculos.csv",
		ModelosFile:  "data/modelos.csv",
	}

	veiculos, errVeiculos := controllers.EstudarVeiculos(csvReader)
	modelos, errModelos := controllers.EstudarModelos(csvReader)

	if errModelos != nil || errVeiculos != nil {
		fmt.Println("Erro:", errModelos, errVeiculos)
		os.Exit(1)
	}

	var saidasEtanol, saidasGasolina []models.Saida

	tempoFilaEtanol, tempoFilaGasolina := 0.0, 0.0

	for _, veiculo := range veiculos {
		modelo := modelos[veiculo.Modelo]
		var combustivel string

		if modelo.ConsumoEtanol > 0 && modelo.CustoPorKmEtanol < modelo.CustoPorKmGasolina {
			combustivel = "ETANOL"

			tempoFilaEtanol += modelo.TempoCarregamentoEtanol
			saida := models.Saida{

				TempoCarregamento: tempoFilaEtanol,
				Menssagem:         controllers.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoGasolina),
			}
			saidasEtanol = append(saidasEtanol, saida)
		} else {
			combustivel = "GASOLINA"

			tempoFilaGasolina += modelo.TempoCarregamentoGasolina
			saida := models.Saida{

				TempoCarregamento: tempoFilaGasolina,
				Menssagem:         controllers.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoGasolina),
			}
			saidasGasolina = append(saidasGasolina, saida)
		}

	}

	controllers.ImprimeSaida(saidasEtanol, saidasGasolina)

}

/*

 */
