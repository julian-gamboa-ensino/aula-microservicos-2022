
#### 05-manipulacao-de-arquivos/main.go
```go
package main

import (
    "fmt"
    "os"
)

func main() {
    arquivo, err := os.Create("exemplo.txt")
    if err != nil {
        fmt.Println("Erro ao criar arquivo:", err)
        return
    }
    defer arquivo.Close()

    _, err = arquivo.WriteString("Olá, Golang!\n")
    if err != nil {
        fmt.Println("Erro ao escrever no arquivo:", err)
    } else {
        fmt.Println("Arquivo escrito com sucesso.")
    }
}

