package main

import "fmt"

func main() {
	s := []int{2, 4, 6, 8, 10}
	fmt.Printf("len = %d, cap = %d e o valor = %d\n", len(s), cap(s), s[:4])
	fmt.Printf("len = %d, cap = %d e o valor = %d", len(s[2:]), cap(s[1:]), s[4])
}