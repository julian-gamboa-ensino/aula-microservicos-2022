package datareader

import "projeto-posto/models"

// DataReader é uma interface para leitura de dados.
type DataReader interface {
    ReadVeiculos() ([]models.Veiculo, error)
    ReadModelos() (map[string]models.Modelo, error)
}
