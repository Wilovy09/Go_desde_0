# Go desde 0

## Como descargar paquetes y herramientas

Para hacer esto en VSCode debemos tener instalada la extensión de Go.

Ejecutamos ``ctrl+shift+p`` o ``f1`` para abrir la paleta de comandos.

Una vez estemos en la paleta de comandos escribiremos `GO: Install/Update Tools` y seleccionaremos todas las herramientas que nos de para actualizar, esperamos a que todas nos regresen un `SUCCEEDED` y listo!. Tenemos nuestro sistema de Go actualizado.

## Hola mundo

Para lograr esto tenemos que crear un archivo con la extensión `.go`.

```go
// En este caso estamos definiendo un paquete main, que es el punto de entrada de nuestra aplicación.
package main

// Importamos el paquete fmt, que nos permite imprimir mensajes en la consola.
import "fmt"

// Definimos la función main, que es el punto de entrada de nuestra aplicación.
func main(){
    /*  Imprimimos un mensaje en la consola.
        usamos el paquete fmt y su metodo Println.
    */
    fmt.Println("Hello, World!")
}
```

Para poder ejecutar nuestro programa usamos el comando `go run`.

```sh
go run NOMBRE.go
```

Esto nos devolvera nuestro `Hello, World!` en nuestra consola.

Pero si queremos compilarlo usaremos `go build`.

```sh
go build NOMBRE.go
```

Esto nos creara un ejecutable que nos permitira mostrar un ``Hello, World!`` en nuestra consola cada que lo ejecutemos.

## Tipos de datos

El lenguaje de programación Go es tipado, es importante conocer los tipos de datos que existen antes de ponernos a trabajar con Go.

- Números enteros
- Números flotantes
- Cadena de textos
- Bolleanos
- Derivados
  - Punteros
  - Arreglos
  - Estructuras
  - Uniones
  - Funciones
  - Slices
  - Interfaces
  - Maps
  - Channels

### Rango de valores de números enteros

| Identificador | Valor                                               |
|---------------|-----------------------------------------------------|
| uint8         | 8bits  (0 a 255)                                    |
| uint16        | 16bits (0 a 65535)                                  |
| uint32        | 32bits (0 a 4294967295)                             |
| uint64        | 64bits (0 a 18446744073709551615)                   |
| int8          | 8bits  (-128 a 127)                                 |
| int16         | 16bits (-32768 a 32767)                             |
| int32         | 32bits (-2147483648 a 2147483647)                   |
| int64         | 64bits (-9223372036854775808 a 9223372036854775807) |

### Rango de valores de números flotantes

- float32
- float64
- complex64
- complex128

### Otra implementación de tipo de datos en Go

| Identificador | Valor                                               |
|---------------|-----------------------------------------------------|
| byte          | 8bits  (0 a 255)                                    |
| rune          | 32bits (-2147483648 a 2147483647)                   |
| uint          | 32bits o 64bits                                     |
| int           | 32bits o 64bits                                     |

```go
package main

import "fmt"

func main(){
    // String
    fmt.Println("Hello, World!")

    // Integer
    fmt.Println(21)

    // Float
    fmt.Println(3.14)

    // Boolean
    fmt.Println(true)
}
```

## Variables

```go
package main

import "fmt"

func main(){
    // Declarar una variable de tipo string
    var nombre string
    // Asignar un valor a la variable
    nombre = "wilovy"

    fmt.Println("Hola, " + nombre)
}
```

Tambien podemos inicializarla con un valor.

```go
package main

import "fmt"

func main(){
    var nombre string = "Wilovy"

    fmt.Println("Hola, " + nombre)
}
```

### Definir multiples variables

En este caso solo definimos multiples variables del mismo tipo.

```go
package main

import "fmt"

func main(){
    var a, b int = 10, 20

    fmt.Println("a = ", a)
    fmt.Println("b = ", b)
}
```

En este caso definimos multiples varibales de distinto tipos.

```go
package main

import "fmt"

func main(){
    var (
        name string = "Wilovy"
        age int = 21

        // Por inferencia
        booleano = true
    )

    fmt.Println("name = ", name)
    fmt.Println("age = ", age)
    fmt.Println("booleano = ", booleano)
}
```

Tambien podemos definir variables tipadas por inferencia sin la necesidad de escribir ``var NOMBRE tipo``.

```go
package main

import "fmt"

func main(){
    // var v1 int = 20
    v1 := 20

    v1 = 21

    fmt.Println(v1)
}
```

En este caso perdemos la opcion de dejar la variable sin ningun valor.

## Constantes

La forma más sencilla de decir que es una constante es:

> Es una variable a la que no se le puede cambiar el valor. Por eso se llaman Contantes, porque su valor no puede cambiar.

```go
package main

import "fmt"

func main() {

    const n int = 21

    // Tambien podemos declarar multiples constantes
    const (
        a string = "Wilovy"
        b bool = true
        c = 3
    )

    fmt.Println("Constante = ", n)

    fmt.Println("Constante a = ", a)
    fmt.Println("Constante b = ", b)
    fmt.Println("Constante c = ", c)
}
```

## Palabras reservadas

Las palabras reservadas no se pueden utilizar como nombres de variables, constantes o de funciones.
<!-- Poner en una tabla -->
- break
- default
- funct
- interface
- select
- case
- defer
- go
- map
- struct
- chan
- else
- goto
- package
- switch
- const
- fallthreough
- if
- range
- type
- continue
- for
- import
- return
- var

## Salida y entrada de datos

Hay distintas formas de mostrar datos en la constola.

```go
package main

import "fmt"

func main() {
    // fmt.Print no hace un salto de línea
    fmt.Print("No salto de línea 1")
    fmt.Print("\tNo salto de línea 2\n")


    // fmt.Println si hace un salto de línea
    fmt.Println("Salto de línea 1")

    // fmt.Printf no hace un salto de línea
    fmt.Printf("No salto de línea 3")
    fmt.Printf("No salto de línea 4")
}
```

- `\n` hace un salto de linea
- `\t` deja un espacio de una tabulación

### Concatenación

Podemos concatenar usando `,` en nuestra salida de datos.

```go
package main

import "fmt"

func main() {
    nombre := "Wilovy"
    edad := 21
    pi := 3.1416
    booleano := true

    fmt.Println("Nombre:", nombre, "\nEdad:", edad, "\nPI:", pi, "\nBooleano:", booleano)
}
```

Tambien podemos "formatear" nuestra salida de datos usando ``fmt.Printf``.

```go
package main

import "fmt"

func main() {
    nombre := "Wilovy"
    edad := 21
    pi := 3.1416
    booleano := true

    fmt.Printf("Nombre: %s \nEdad: %d \nPi: %f \nBooleano: %t", nombre, edad, pi, booleano)
}
```

- ``%s`` sirve para concatenar strings
- ``%d`` sirve para concatenar números enteros
- ``%f`` sirve para concatenar números flotantes
- ``%t`` sirve para concatenar booleanos

### Entrada de datos

A veces necesitamos que el usuario nos mande el dato que requerimos para poder continuar con la ejecución de nuestro programa.

```go
package main

import "fmt"

func main() {
    var (
        nombre string
        edad int
        pi float64
        booleano bool
    )

    fmt.Print("Ingresa tu nombre: ")
    /*  Con Scanln indicamos en que variable va a guardar el input del usuario.
        Y dejamos de capturar el input cuando el usuario presiona la tecla Enter o hay un salto de linea.
    */
    fmt.Scanln(&nombre)

    fmt.Print("Ingresa tu edad: ")
    fmt.Scanln(&edad)

    fmt.Print("Ingresa el valor de pi: ")
    fmt.Scanln(&pi)

    fmt.Print("Ingresa un valor booleano: ")
    fmt.Scanln(&booleano)

    fmt.Printf("Nombre: %s \nEdad: %d \nPi: %f \nBooleano: %t", nombre, edad, pi, booleano)
}
```

- ``Scanln`` indicamos en que variable va a guardar el input del usuario y dejamos de capturar el input cuando el usuario presiona la tecla Enter o hay un salto de linea.
- ``Scanf`` indicamos el tipo de dato que queremos capturar (``%s``, ``%d``, ``%f``, ``%t``, etc...) y no deja de capturar hasta que haya un salto de linea.

```go
    fmt.Print("Ingresa tu edad: ")
    fmt.Scanf("%d", &edad)
```

## Operadores aritméticos

Con operadores atitméticos podemos realizar operaciones matemáticas ya sea con valores literales o con variables. a = 10 y b = 5

| Operador | Descripción                    | Ejemplo        |
|----------|--------------------------------|----------------|
| +        | Sumar                          | a+b resulta 15 |
| -        | Restar                         | a-b resulta 5  |
| *        | Multiplicar                    | a*b resulta 50 |
| /        | Dividir                        | a/b resulta 2  |
| %        | Calcular módulo de la división | a%b resulta 0  |

### Ejericicio #1

Crear una aplicación que pide el valor de ``a``, el valor de ``b`` y hacer todas las operaciones mencionadas con esos 2 números proporcionados por el usuario.

#### Respuesta ejercicio #1

```go
package main

import "fmt"

func main() {
    var a,b int

    fmt.Print("Valor de a: ")
    fmt.Scanln(&a)

    fmt.Print("Valor de b: ")
    fmt.Scanln(&b)

    fmt.Println("Resultado de la suma: ", a + b)
    fmt.Println("Resultado de la resta: ", a - b)
    fmt.Println("Resultado de la multiplicación: ", a * b)
    fmt.Println("Resultado de la división: ", a / b)
    fmt.Println("Resultado del módulo: ", a % b)
}
```

## Operadores relacionales

Los operadores relacionales lo que hacen es comparar 2 valores y devuelven un valor booleano de acuerdo a la compración si es verdad o falso. a = 1 y b = 2

| Operador | Descripción       | Ejemplo                    |
|----------|-------------------|----------------------------|
| ==       | Igualdad          | a == b devuelve falso      |
| !=       | Distintos         | a != b devuelve verdadero  |
| >        | Mayor que         | a > b devuelve falso       |
| <        | Menor que         | a < b devuelve falso       |
| >=       | Mayor o igual que | a >= b devuelve falso      |
| <=       | Menor o igual que | a <= b devuelve verdadero  |

### Ejercicio #2

Crear una aplicación que pide el valor de ``a``, el valor de ``b`` y hacer todas las operaciones mencionadas con esos 2 números proporcionados por el usuario.

#### Respuesta ejercicio #2

```go
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
```

## Condicionales

```go
package main

import "fmt"

func main() {

    // Si es false haz:
    if false {
        fmt.Println("Se cumple la condición")
    }
    // si no es falso haz:
    else {
        fmt.Println("No se cumple la condición")
    }
}
```

### Ejercicio #3

Crear una aplicación que pide el valor de ``a``, el valor de ``b`` e imprimir si es par o impar.

#### Respuesta ejercicio #3

```go
package main

import "fmt"

func main() {

    var a int

    fmt.Println("Ingrese un número: ")
    fmt.Scan(&a)

    if a%2 == 0{
        fmt.Println("Es par")
    } else {
        fmt.Println("Es impar")
    }
}
```

Y si tenemos mas de una condición?, para eso usamos ``else if``

```go
package main

import "fmt"

func main() {

    var a int

    fmt.Println("Ingrese un número: ")
    fmt.Scan(&a)

    // si a es igual a 0 haz:
    if a == 0{
        fmt.Println("Es neutro")
    }
    // si el módulo de a%2 es 0 haz:
    else if a%2 == 0{
        fmt.Println("Es par")
    }
    // si no se cumple ninguna haz:
    else {
        fmt.Println("Es impar")
    }
}
```

## Operadores lógicos

Los operadores lógicos lo que hace es comprarar dos valores booleanos, tenemos 3 operadores:

- **NOT !**, lo que hace negar su valor. Si su valor es true lo negara a false.

```go
package main

import "fmt"

func main() {

    fmt.Println(! true)
}
```

- **AND &&**, lo que hace es comprara 2 valores booleanos y devuelve otro valor booleano, solo cuanbdo ambos son true devuelve true.

```go
package main

import "fmt"

func main() {

    fmt.Println(true && true)
}
```

- **OR ||**, lo que hace es también comparar 2 valores booleanos y devuelve otro valore booleano en este caso ambos o al menos uno tiene que ser true para que devuelva true.

```go
package main

import "fmt"

func main() {

    fmt.Println(true || false)
}
```

## Switch ó Casos

Los casos trabajan dentro de la instrucción switch, son estructuras selectiva múltiple donde puede realizar varios casos.

```go
package main

import "fmt"

func main() {

    var n int

    fmt.Println("1 o 2:")
    fmt.Scanln(&n)

    switch n {
    // Si n == 1 haz:
    case 1:
        fmt.Println("El número es 1")
    // Si n == 2 haz:
    case 2:
        fmt.Println("El número es 2")
    // Si n no coincide con ningún caso haz:
    default:
        fmt.Println("El número ingresado no es 1 ni 2")
    }
}
```

### Ejercicio #4

Realiza un sistema que pida ingresar un número del 1 al 5 y devuelva escrito como por ejemplo, ingrese un número que sea 3 y devuelve TRES

```go
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
```

## Operadores en asignación

a = 5

| Operador | Descripción                  | Ejemplo          |
|----------|------------------------------|------------------|
| =        | Asignar un valor             | a = 5            |
| +=       | Suma en asignación           | a = a+3  -> a+=3 |
| -=       | Resta en asignación          | a -= 3           |
| *=       | Multiplicación en asignación | a *= 3           |
| /=       | División en asignación       | a /= 3           |
| %=       | Módulo en asignación         | a %= 3           |
| ++       | Suma 1                       | a++              |
| --       | Resta 1                      | a--              |

## Bucle For

El Go no existe ``while``, ``do while`` simplemente existe el bucle ``For`` y puede ser usado como ``while``, ``do while``.

```go
package main

import "fmt"

func main() {

    // como for
    for i := 0; i < 5; i++{
        fmt.Println(i)
    }

    // como do while
    var c int = 0
    for c <= 5{
        fmt.Println(c)
        c++
    }

    // como while
    var a int = 0
    for true{
        fmt.Println(a)
        a++
    }
}
```

### Break y Continue

Las palabras reservadas ``break`` y ``continue``. Sirven para detener o continuar la ejecución de nuestro bucle.

```go
package main

import "fmt"

func main() {

    for i := 0; i < 10; i++{

        if i == 8 {
            fmt.Println("break")
            break
        } else {
            fmt.Println("Continue", i)
            continue
        }

    }

}
```

## Manejo de datos

### Array / Arreglos

Array es un almacenador de datos fijos, puede almacenar cantidad de datos pero indicando cuantos datos va almacenar ese array, y en Go, como es tipado solo puede almacenar un tipo de datos.

- Un array se define como las vanables solo que en array se coloca dentro de los corchetes la cantidad de datos que va almacenar y para agregar y modificar su datos usamos el índice de cada dato que se coloca entre corchetes.

```go
package main

import "fmt"

func main() {

    // Declararar un array de 2 elementos
    var array1 [2]string

    // Asignar valores a los elementos del array
    array1[0] = "Hello"
    array1[1] = "World"

    // Imprimir el array
    fmt.Println(array1)

    // Declarar e inicializar un array de 2 elementos
    array2 := [2]string{"Hello", "World"}
    // Imprimir el array
    fmt.Println(array2)
}
```

### Slicen

Los slicen son parecidas a los array que también almacena datos pero los slicen almacena cantidad de datos indeterminados, que puede almacenar cantidad de vanables que tu desees.

- Puedes definir un slicen vacio y luego agregar los datos con la función append.

```go
package main

import "fmt"

func main() {

    // Creamos un slice de strings
    var slicen1 []string

    // Agregamos elementos al slice
    slicen1 = append(slicen1, "a")
    slicen1 = append(slicen1, "b", "c", "d")

    // Imprimimos el slice1
    fmt.Println(slicen1)

    // Creamos un slice de strings y lo inicializamos con valores
    slicen2 := []string{"e", "f"}
    fmt.Println(slicen2)
}
```

## Funciones

Una función es útil para ordenar nuestros códigos, para reutilizar códigos, para hacer más legible nuestro código.

- Una función puede recibir parámetros y devolver un valor.
- Una función se puede definir con la palabra reservada ``func`` seguido del nombre de la función y los parámetros que va a recibir y el tipo de dato que va a devolver.
- Para llamar a una función solo se escribe el nombre de la función y los parámetros que va a recibir.

```go
package main

import "fmt"

// Definir una función con parámetros y retorno
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
```

---

## Go

```sh
# Para inicializar un modulo de go
go mod init NOMBRE
```

```sh
# Para limpiar los modulos que no se estan usando
go mod tidy
```

```sh
# Para descargar los modulos que se estan usando
go mod vendor
```

```sh
# Para ver los modulos que se estan usando
go list -m all

go list -m -json all
```

```sh
# Para ver la documentación de un paquete
go doc fmt
```
