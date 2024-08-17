
#### 07-testes/main.go
```go
package main

import "fmt"

func Soma(a, b int) int {
    return a + b
}

func main() {
    resultado := Soma(5, 7)
    fmt.Println("Resultado da soma:", resultado)
}

