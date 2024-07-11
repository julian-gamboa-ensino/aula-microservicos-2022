package main

import (
	"fmt"
	"os"
	"projeto-posto/controller"
	"projeto-posto/datareader"
	"projeto-posto/service"
)

func main() {
	csvReader := &datareader.CSVReader{
		VeiculosFile: "data/veiculos.csv",
		ModelosFile:  "data/modelos.csv",
	}

	veiculos, errVeiculos := controller.EstudarVeiculos(csvReader)
	modelos, errModelos := controller.EstudarModelos(csvReader)

	if errModelos != nil || errVeiculos != nil {
		fmt.Println("Erro:", errModelos, errVeiculos)
		os.Exit(1)
	}
	postoService := service.NewPostoService()
	saidasEtanol, saidasGasolina := postoService.ProcessarVeiculos(veiculos, modelos)

	controller.ImprimeSaida(saidasEtanol, saidasGasolina)

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
