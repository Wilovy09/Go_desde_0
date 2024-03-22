package main

import "fmt"

func main() {

    var n int

    fmt.Println("Ingrese un número del 1 al 5:")
    fmt.Scanln(&n)

    switch n {
    case 1:
        fmt.Println("UNO")
    case 2:
        fmt.Println("DOS")
    case 3:
        fmt.Println("TRES")
    case 4:
        fmt.Println("CUATRO")
    case 5:
        fmt.Println("CINCO")
    default:
        fmt.Println("El número ingresado no está en el rango de 1 al 5")
    }
}