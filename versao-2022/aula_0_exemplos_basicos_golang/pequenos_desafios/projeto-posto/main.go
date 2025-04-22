package main

import (
	"fmt"
	"os"
	"projeto-posto/controller"
	"projeto-posto/datareader"
	"projeto-posto/service"
)

func main() {
	csvReader := &datareader.CSVReader{
		VeiculosFile: "data/veiculos.csv",
		ModelosFile:  "data/modelos.csv",
	}

	veiculos, errVeiculos := controller.EstudarVeiculos(csvReader)
	modelos, errModelos := controller.EstudarModelos(csvReader)

	if errModelos != nil || errVeiculos != nil {
		fmt.Println("Erro:", errModelos, errVeiculos)
		os.Exit(1)
	}
	postoService := service.NewPostoService()
	saidasEtanol, saidasGasolina := postoService.ProcessarVeiculos(veiculos, modelos)

	controller.ImprimeSaida(saidasEtanol, saidasGasolina)

}
