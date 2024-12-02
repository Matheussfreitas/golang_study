package main

import (
    "fmt"
)

func main() {
	// Memória -> Endereço -> Valor
	a := 10
	var ponteiro *int = &a
	*ponteiro = 20
	fmt.Println("Endereço de a:", &ponteiro)
}