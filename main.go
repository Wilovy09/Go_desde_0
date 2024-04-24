package main

import (
    "fmt"
)

func main() {

    colores := map[string]string{
        "rojo":   "#ff0000",
        "verde":  "#4bf745",
        "azul":   "#0000ff",
    }

    for clave, valor := range colores {
        fmt.Println(clave, "=>", valor)
    }
}