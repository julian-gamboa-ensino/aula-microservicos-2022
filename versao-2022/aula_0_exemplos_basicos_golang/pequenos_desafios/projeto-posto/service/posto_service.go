package service

import (
	"projeto-posto/controller"
	"projeto-posto/model"
)

type PostoService interface {
	ProcessarVeiculos(
		veiculos []model.Veiculo,
		modelos map[string]model.Modelo) ([]model.Saida, []model.Saida)
}

type postoService struct{}

func NewPostoService() PostoService {
	return &postoService{}
}

func (ps *postoService) ProcessarVeiculos(veiculos []model.Veiculo, modelos map[string]model.Modelo) ([]model.Saida, []model.Saida) {
	var saidasEtanol, saidasGasolina []model.Saida
	tempoFilaEtanol, tempoFilaGasolina := 0.0, 0.0

	for _, veiculo := range veiculos {
		modelo := modelos[veiculo.Modelo]
		var combustivel string

		if modelo.ConsumoEtanol > 0 && modelo.CustoPorKmEtanol < modelo.CustoPorKmGasolina {
			combustivel = "ETANOL"
			tempoFilaEtanol += modelo.TempoCarregamentoEtanol
			saida := model.Saida{
				TempoCarregamento: tempoFilaEtanol,
				Menssagem:         controller.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoEtanol),
			}
			saidasEtanol = append(saidasEtanol, saida)
		} else {
			combustivel = "GASOLINA"
			tempoFilaGasolina += modelo.TempoCarregamentoGasolina
			saida := model.Saida{
				TempoCarregamento: tempoFilaGasolina,
				Menssagem:         controller.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoGasolina),
			}
			saidasGasolina = append(saidasGasolina, saida)
		}
	}

	return saidasEtanol, saidasGasolina
}

/*


func processarVeiculo(veiculo models.Veiculo, modelos map[string]models.Modelo, saidasEtanol, saidasGasolina *[]models.Saida, tempoFilaEtanol, tempoFilaGasolina *float64, wg *sync.WaitGroup) {
    defer wg.Done() // Sinaliza que a goroutine terminou

    modelo := modelos[veiculo.Modelo]
    var combustivel string

    if modelo.ConsumoEtanol > 0 && modelo.CustoPorKmEtanol < modelo.CustoPorKmGasolina {
        combustivel = "ETANOL"
        *tempoFilaEtanol += modelo.TempoCarregamentoEtanol // Acesso sincronizado
        saida := models.Saida{
            TempoCarregamento: *tempoFilaEtanol,
            Menssagem:        controllers.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoGasolina),
        }
        *saidasEtanol = append(*saidasEtanol, saida) // Acesso sincronizado
    } else {
        combustivel = "GASOLINA"
        *tempoFilaGasolina += modelo.TempoCarregamentoGasolina // Acesso sincronizado
        saida := models.Saida{
            TempoCarregamento: *tempoFilaGasolina,
            Menssagem:        controllers.PreparaSaida(modelo, veiculo, combustivel, modelo.TempoCarregamentoGasolina),
        }
        *saidasGasolina = append(*saidasGasolina, saida) // Acesso sincronizado
    }
}

func main() {
    // ... (Leitura de dados igual ao código original) ...

    var saidasEtanol, saidasGasolina []models.Saida
    tempoFilaEtanol, tempoFilaGasolina := 0.0, 0.0

    var wg sync.WaitGroup // WaitGroup para sincronizar as goroutines
    wg.Add(len(veiculos)) // Indica quantas goroutines serão iniciadas

    for _, veiculo := range veiculos {
        go processarVeiculo(veiculo, modelos, &saidasEtanol, &saidasGasolina, &tempoFilaEtanol, &tempoFilaGasolina, &wg)
    }

    wg.Wait() // Espera todas as goroutines terminarem

    controllers.ImprimeSaida(saidasEtanol, saidasGasolina)
}
*/
