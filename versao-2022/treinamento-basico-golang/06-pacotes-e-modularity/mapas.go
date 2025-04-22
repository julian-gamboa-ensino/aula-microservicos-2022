package main

import "fmt"

func main() {
    capitais := map[string]string{
        "São Paulo": "São Paulo",
        "Rio de Janeiro": "Rio de Janeiro",
        "Minas Gerais": "Belo Horizonte",
    }

    for estado, capital := range capitais {
        fmt.Printf("A capital do estado %s é %s\n", estado, capital)
    }
}