package main

import "fmt"

func main() {
	fmt.Println(sum(7,3))
	fmt.Println(somar(12,566,18,23,4,841,7,3))

	// funções closures ou anônimas
	total := func() int {
		return sum(1,2) * somar(11,12,13,14,15)
	}()
	fmt.Println(total)
}

func sum(a, b int) int {
	return a + b
}

// funções variádicas, quando não se sabe a quantidade de parametros
func somar(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
        total += numero
    }
	return total
}