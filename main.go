package main

import "fmt"

// Definir una función con parámetros y retorno
// La función suma recibe dos parámetros de tipo entero y retorna un entero
func suma(a int, b int) int {
    return a + b
}

// Nuestra función principal
func main() {

    // Llamar a la función y mostrar el resultado
    fmt.Println(suma(5, 5))

    // Llamar a la función saludo
    saludo()
}

// Definir una función sin parámetros y sin retorno
func saludo() {
    fmt.Println("Hola, Wilovy")
}