package main

import (
    "fmt"
)

func main() {

    // Creamos una nueva carrera vacia new(Carrera)
    carrera1 := new(Carrera)

    carrera1.nombreCarrera = "Programación backend"
    carrera1.duracionCarrera = 5
    carrera1.nombre = "Curso de Go"
    carrera1.url = "https://www.google.com"
    carrera1.habilidades = []string{"Go", "Backend"}

    carrera1.inscribirse("Wilovy")

    fmt.Println(carrera1)
}
// Curso nueva estructura
type Curso struct {
    nombre string
    url string
    habilidades []string
}
// inscribirse es un método de la estructura Curso
func (c Curso) inscribirse(nombre string){
    fmt.Printf("La persona %s se ha inscrito al curso %s \n", nombre, c.nombre)
}

// Carrera nueva estructura
type Carrera struct {
    nombreCarrera string
    duracionCarrera int
    Curso
}