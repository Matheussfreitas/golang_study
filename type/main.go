package main

import (
	"fmt"
)

type ID int

var (
	a string = "abacaxi"
	b bool   = true
	c ID = 1
)

func main() {
	fmt.Printf("O tipo de C é %T", c)
}
