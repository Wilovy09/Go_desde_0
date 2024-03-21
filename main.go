package main

import "fmt"

func main() {
    var a,b int

    fmt.Print("Valor de a: ")
    fmt.Scanln(&a)

    fmt.Print("Valor de b: ")
    fmt.Scanln(&b)

    fmt.Println("Resultado de a igual a b: ", a == b)
    fmt.Println("Resultado de a diferente de b: ", a != b)
    fmt.Println("Resultado de a mayor que b: ", a > b)
    fmt.Println("Resultado de a menor que b: ", a < b)
    fmt.Println("Resultado de a mayor o igual que b: ", a >= b)
    fmt.Println("Resultado de a menor o igual que b: ", a <= b)
}