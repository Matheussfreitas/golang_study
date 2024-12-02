package main

import "fmt"

func main() {
    var meuArray [3]int
    meuArray[0] = 10
    meuArray[1] = 20
    meuArray[2] = 50
    fmt.Println(meuArray)
    fmt.Println(len(meuArray))
    fmt.Println(meuArray[2])

    for  i, v := range meuArray {
        fmt.Printf("O valor do inidce é %d e o valor é %d\n", i, v)
    }
}