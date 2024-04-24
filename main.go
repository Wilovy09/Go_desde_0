package main

import (
    "fmt"
)

func main() {
    // Crear una instancia de la estructura Curso y asignar valores a sus campos
    curso1 := Curso{
        nombre: "Curso profesional de Go",
        url: "https://go.com/",
        habilidades: []string{"Go", "Backend"},
    }

    fmt.Println(curso1)
    fmt.Println(curso1.nombre)
    fmt.Println(curso1.url)
    fmt.Println(curso1.habilidades)

    // Crear una instancia de la estructura Curso vacia
    curso2 := new(Curso)

    // Asignar valores a los campos de la estructura Curso
    curso2.nombre = "Curso de Python"
    curso2.url = "https://python.com/"
    curso2.habilidades = []string{"Python", "Backend"}

    fmt.Println(curso2)
    fmt.Println(curso2.nombre)
    fmt.Println(curso2.url)
    fmt.Println(curso2.habilidades)
}

// Curso es una estructura que representa un curso
type Curso struct {
    nombre string
    url string
    habilidades []string
}