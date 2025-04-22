package controller

import (
	"projeto-posto/datareader"
	"projeto-posto/model"
)

func EstudarVeiculos(dr datareader.DataReader) ([]model.Veiculo, error) {
	veiculos, err := dr.ReadVeiculos()
	if err != nil {
		return nil, err
	}
	return veiculos, nil
}
