package controllers

import (
	"projeto-posto/datareader"
	"projeto-posto/models"
)

func EstudarModelos(dr datareader.DataReader) (map[string]models.Modelo, error) {
	modelos, err := dr.ReadModelos()
	if err != nil {
		return nil, err
	}
	return modelos, nil
}
