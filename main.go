package main

import (
    "fmt"
    "github.com/Wilovy09/Go_desde_0/mensaje"
)

func main() {
    fmt.Println("Hola desde main.go")
    mensaje.SaludoPublico()
    // mensaje.saludoPrivado() // Error: saludoPrivado is not visible

    mensaje.OtraFuncion()
}
