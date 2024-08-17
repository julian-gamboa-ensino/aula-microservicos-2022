
#### 06-pacotes-e-modularity/main.go
```go
package main

import (
    "fmt"
    "github.com/usuario/meu-pacote"
)

func main() {
    mensagem := meu_pacote.Saudacao("Mundo")
    fmt.Println(mensagem)
}

