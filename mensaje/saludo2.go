package mensaje

import "fmt"

// Ejecuta una fución privada de saludo1.go
func OtraFuncion() {
	fmt.Println("\nEjecutando la función privada desde otra función en otro archivo del mismo paquete")

    /* llamamos a la saludoPrivado directamente, ya que si estamos en el mismo paquete
     * podemos llamar a las funciones de otros archivos
     */
    saludoPrivado()
}