package controller

import (
	"projeto-posto/datareader"
	"projeto-posto/model"
)

func EstudarModelos(dr datareader.DataReader) (map[string]model.Modelo, error) {
	modelos, err := dr.ReadModelos()
	if err != nil {
		return nil, err
	}
	return modelos, nil
}
