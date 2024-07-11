package datareader

import (
	"encoding/csv"
	"fmt"
	"os"
	"projeto-posto/model"

	"strconv"
	"strings"
)

type CSVReader struct {
	VeiculosFile string
	ModelosFile  string
}

func (r *CSVReader) ReadVeiculos() ([]model.Veiculo, error) {
	file, err := os.Open(r.VeiculosFile)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo %s: %v", r.VeiculosFile, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo %s: %v", r.VeiculosFile, err)
	}

	var veiculos []model.Veiculo
	for _, record := range records[1:] {
		veiculo := model.Veiculo{
			Modelo: record[0],
			Placa:  record[1],
		}
		veiculos = append(veiculos, veiculo)
	}
	return veiculos, nil
}

func (r *CSVReader) ReadModelos() (map[string]model.Modelo, error) {
	file, err := os.Open(r.ModelosFile)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir o arquivo %s: %v", r.ModelosFile, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ';'
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo %s: %v", r.ModelosFile, err)
	}

	modelos := make(map[string]model.Modelo)
	for i, record := range records[1:] {
		if len(record) < 4 {
			return nil, fmt.Errorf("linha %d do arquivo %s tem formato inválido", i+2, r.ModelosFile)
		}

		nome := record[0]
		consumoEtanolStr := strings.ReplaceAll(record[1], ",", ".")
		consumoGasolinaStr := strings.ReplaceAll(record[2], ",", ".")
		capacidadeTanqueStr := strings.ReplaceAll(record[3], ",", ".")

		consumoEtanol, err := strconv.ParseFloat(consumoEtanolStr, 64)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter consumo de etanol na linha %d: %v", i+2, err)
		}
		consumoGasolina, err := strconv.ParseFloat(consumoGasolinaStr, 64)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter consumo de gasolina na linha %d: %v", i+2, err)
		}
		capacidadeTanque, err := strconv.ParseFloat(capacidadeTanqueStr, 64)
		if err != nil {
			return nil, fmt.Errorf("erro ao converter capacidade do tanque na linha %d: %v", i+2, err)
		}

		custoPorKmEtanol := 2.27 / consumoEtanol
		custoPorKmGasolina := 2.90 / consumoGasolina

		vazamentoBombaEtanol := 12.0
		vazamentoBombaGasolina := 10.0

		tempoCarregamentoEtanol := capacidadeTanque / vazamentoBombaEtanol
		tempoCarregamentoGasolina := capacidadeTanque / vazamentoBombaGasolina

		modelos[nome] = model.Modelo{
			Nome:                      nome,
			ConsumoEtanol:             consumoEtanol,
			ConsumoGasolina:           consumoGasolina,
			CapacidadeTanque:          capacidadeTanque,
			CustoPorKmEtanol:          custoPorKmEtanol,
			CustoPorKmGasolina:        custoPorKmGasolina,
			TempoCarregamentoGasolina: tempoCarregamentoGasolina,
			TempoCarregamentoEtanol:   tempoCarregamentoEtanol,
		}
	}
	return modelos, nil
}
