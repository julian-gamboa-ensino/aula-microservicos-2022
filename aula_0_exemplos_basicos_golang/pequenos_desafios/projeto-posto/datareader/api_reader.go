package datareader

import (
	"encoding/json"
	"net/http"
	"projeto-posto/model"
)

type APIDataReader struct {
	VeiculosURL string
	ModelosURL  string
}

func (r *APIDataReader) ReadVeiculos() ([]model.Veiculo, error) {
	resp, err := http.Get(r.VeiculosURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var veiculos []model.Veiculo
	if err := json.NewDecoder(resp.Body).Decode(&veiculos); err != nil {
		return nil, err
	}

	return veiculos, nil
}

func (r *APIDataReader) ReadModelos() (map[string]model.Modelo, error) {
	resp, err := http.Get(r.ModelosURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var modelos map[string]model.Modelo
	if err := json.NewDecoder(resp.Body).Decode(&modelos); err != nil {
		return nil, err
	}

	return modelos, nil
}
