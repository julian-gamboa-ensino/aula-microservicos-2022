package datareader

import "projeto-posto/model"

// DataReader é uma interface para leitura de dados.
type DataReader interface {
	ReadVeiculos() ([]model.Veiculo, error)
	ReadModelos() (map[string]model.Modelo, error)
}
