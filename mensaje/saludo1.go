package mensaje

import "fmt"

// saludoPrivado es un mensaje privado
func saludoPrivado() {
    fmt.Println("Hola desde mensaje/saludo1.go (saludoPrivado)")
}

// SaludoPublico es un mensaje público
func SaludoPublico() {
    fmt.Println("Hola desde mensaje/saludo1.go (SaludoPublico)")
}