package main

import "fmt" 

func intSeq() func() int {
    i := 0
    return func() int {
        i += 1
        return i 
    }
}

func main() {
    nextInt := intSeq() // Aqui "nextInt" recebe como valor a função intSeq
    fmt.Println(nextInt()) // Printa 1
    fmt.Println(nextInt()) // Printa 2
    fmt.Println(nextInt()) // Printa 3
    newInts := intSeq()
    fmt.Println(newInts()) // Printa 1 
}