package controllers

import (
	"projeto-posto/datareader"
	"projeto-posto/models"
)

func EstudarVeiculos(dr datareader.DataReader) ([]models.Veiculo, error) {
	veiculos, err := dr.ReadVeiculos()
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}
