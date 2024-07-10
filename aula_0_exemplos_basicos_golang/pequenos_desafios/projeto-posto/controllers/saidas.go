package controllers

import (
	"fmt"
	"projeto-posto/models"
	"sort"
)

// Função para ordenar um slice de Saida por TempoCarregamento
type ByTempoCarregamento []models.Saida

func (a ByTempoCarregamento) Len() int      { return len(a) }
func (a ByTempoCarregamento) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByTempoCarregamento) Less(i, j int) bool {
	return a[i].TempoCarregamento < a[j].TempoCarregamento
}

func formatMinutes(minutes float64) string {
	minPart := int(minutes)                           // Parte inteira dos minutos
	secPart := int((minutes - float64(minPart)) * 60) // Parte decimal convertida para segundos
	return fmt.Sprintf("%02d:%02d", minPart, secPart)
}

func PreparaSaida(modelo models.Modelo, veiculo models.Veiculo, combustivel string, tempoCarregamento float64) string {
	return fmt.Sprintf("Veículo modelo %s, placa %s foi abastecido com %.2f litros de %s .\n",
		veiculo.Modelo, veiculo.Placa, modelo.CapacidadeTanque, combustivel)
}

func ImprimeSaida(saidas_etanol []models.Saida, saidas_gasolina []models.Saida) {
	todasSaidas := append(saidas_gasolina, saidas_etanol...)
	sort.Sort(ByTempoCarregamento(todasSaidas))

	for _, saida := range todasSaidas {
		fmt.Printf("[%s] %s ", formatMinutes(saida.TempoCarregamento), saida.Menssagem)
	}
}
