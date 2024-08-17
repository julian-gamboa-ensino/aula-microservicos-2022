package main

import (
	"fmt"
	"time"
)

func digaOla() {
	fmt.Println("Olá!")
}

func main() {
	go digaOla()
	time.Sleep(1 * time.Second)
	fmt.Println("Programa finalizado.")
}
