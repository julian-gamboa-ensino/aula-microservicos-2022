## Projeto Posto de Gasolina (Simulação)

Este projeto em Go simula o funcionamento de um posto de gasolina, com foco na gestão de filas e abastecimento de veículos. Ele lê dados de veículos e modelos de carros de arquivos CSV, calcula os custos e tempos de abastecimento, e imprime um relatório do processo.

**Estrutura do Projeto**

* **`controllers/`:** Contém a lógica principal do programa:
    * **`veiculos.go`:** Funções para ler e processar dados de veículos.
    * **`modelos.go`:** Funções para ler e processar dados de modelos de carros.
    * **`saidas.go`:** Funções para preparar e imprimir as saídas do programa.
* **`datareader/`:** Lida com a leitura de dados de arquivos CSV:
    * **`datareader.go`:** Define a interface `DataReader`.
    * **`csvreader.go`:** Implementa a interface `DataReader` para ler dados de CSV.
* **`models/`:** Define as estruturas de dados:
    * **`veiculo.go`:** Estrutura para representar um veículo.
    * **`modelo.go`:** Estrutura para representar um modelo de carro.
    * **`saida.go`:** Estrutura para representar a saída do programa (abastecimento).
* **`main.go`:** Ponto de entrada do programa.

**Como Funciona**

1. **Leitura de Dados:** Os dados de veículos (`veiculos.csv`) e modelos (`modelos.csv`) são lidos dos arquivos CSV.
2. **Cálculo de Custos:** O programa calcula os custos por quilômetro para etanol e gasolina com base no consumo e preço dos combustíveis.
3. **Determinação do Combustível:** Para cada veículo, o programa determina qual combustível (etanol ou gasolina) é mais vantajoso com base no custo por quilômetro.
4. **Cálculo do Tempo de Abastecimento:** O tempo de abastecimento é calculado com base na capacidade do tanque e na taxa de vazão da bomba.
5. **Simulação da Fila:** Os veículos são adicionados às filas de etanol ou gasolina, e o tempo de espera em cada fila é atualizado.
6. **Impressão da Saída:** O programa imprime um relatório mostrando a ordem de abastecimento dos veículos, o tempo de espera na fila e o combustível utilizado.

**Como Executar**

1. **Pré-requisitos:** Certifique-se de ter o Go instalado em seu sistema.
2. **Dados:** Coloque os arquivos `veiculos.csv` e `modelos.csv` na pasta `data/`.
3. **Compilação:** Execute o comando `go build` para compilar o projeto.
4. **Execução:** Execute o comando `./projeto-posto` para iniciar a simulação.

**Observações**

* O programa assume que os arquivos CSV estão no formato correto e que os dados são válidos.
* Os cálculos de custo e tempo de abastecimento são simplificados para fins de simulação.
* A simulação não leva em conta fatores como a disponibilidade das bombas ou a ordem de chegada dos veículos.

**Melhorias Futuras**

* Adicionar mais realismo à simulação, considerando fatores como a ordem de chegada, a disponibilidade das bombas e a variação no tempo de abastecimento.
* Implementar uma interface gráfica para visualizar a simulação.
* Permitir que o usuário interaja com a simulação, adicionando ou removendo veículos e alterando os parâmetros.

**Contribuições**

Contribuições são bem-vindas! Sinta-se à vontade para abrir issues ou enviar pull requests. 