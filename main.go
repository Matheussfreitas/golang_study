package main


var (
	a string = "abacaxi"
	b bool   = true
)

func main() {
	//short declaration, o tipo é inferido
	c := 10
	println(c, a, b)
}
