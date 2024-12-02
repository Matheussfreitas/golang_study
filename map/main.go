package main

import "fmt"

func main() {
	salarios := map[string]int{"matheus": 800, "mari": 1000,}
	fmt.Println(salarios, "matheus")
	delete(salarios, "matheus")

	salario := make(map[string]int)
	salario["mari"] = 1000
}