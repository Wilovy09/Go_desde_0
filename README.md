# Go desde 0

- Go es un lenguaje de programación de código abierto que facilita la creación de software simple, confiable y eficiente.
- Go es un lenguaje de programación compilado y concurrente, desarrollado por Google.

* [ ] Como dockerizar una app sencilla (sin paquetes externos)
* [ ] Como dockerizar una app comnpleja (con paquetes externos)

## Instalar Go

Para instalar Go en tu sistema operativo, puedes seguir los siguientes pasos:

### Windows

1. Descarga el instalador de Go desde la [página oficial](https://go.dev/).
2. Ejecuta el instalador y sigue los pasos que te indica.
3. Abre una terminal y escribe `go version` para verificar que la instalación fue exitosa.
4. Crea una carpeta en tu escritorio llamada `Go_desde_0`.
5. Abre tu editor de código favorito y abre la carpeta `Go_desde_0`.
6. Crea un archivo llamado `main.go`.
7. Dentro de `main.go` escribe el siguiente código:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

8. Guarda el archivo y en la terminal escribe `go run main.go`.
9. Deberías ver el mensaje `Hello, World!` en la terminal.

### Linux

1. Abre una terminal y escribe el siguiente comando:

```sh
sudo apt-get update
```

2. Luego, escribe el siguiente comando:

```sh
sudo apt-get install golang
```

3. Abre una terminal y escribe `go version` para verificar que la instalación fue exitosa.
4. Crea una carpeta en tu escritorio llamada `Go_desde_0`.
5. Abre tu editor de código favorito y abre la carpeta `Go_desde_0`.
6. Crea un archivo llamado `main.go`.
7. Dentro de `main.go` escribe el siguiente código:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

8. Guarda el archivo y en la terminal escribe `go run main.go`.
9. Deberías ver el mensaje `Hello, World!` en la terminal.

### MacOS

1. Abre una terminal y escribe el siguiente comando:

```sh
brew install go
```

2. Abre una terminal y escribe `go version` para verificar que la instalación fue exitosa.
3. Crea una carpeta en tu escritorio llamada `Go_desde_0`.
4. Abre tu editor de código favorito y abre la carpeta `Go_desde_0`.
5. Crea un archivo llamado `main.go`.
6. Dentro de `main.go` escribe el siguiente código:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

7. Guarda el archivo y en la terminal escribe `go run main.go`.
8. Deberías ver el mensaje `Hello, World!` en la terminal.

## Como descargar paquetes y herramientas

Para hacer esto en [VSCode](https://code.visualstudio.com/) debemos tener instalada la [extensión de Go](https://marketplace.visualstudio.com/items?itemName=golang.go).

Ejecutamos ``ctrl+shift+p`` o ``f1`` para abrir la paleta de comandos.

Una vez estemos en la paleta de comandos escribiremos `GO: Install/Update Tools` y seleccionaremos todas las herramientas que nos de para actualizar, esperamos a que todas nos regresen un `SUCCEEDED` en el apartado SALIDA de nuestra terminal. Tendremos nuestro sistema de Go actualizado.

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

## Operaciones con Cadenas

Para hacer operaciones con cadenas de caracteres tenemos que importar otra libereria que se llama [strings](https://pkg.go.dev/strings).

```go
import (
    "fmt"
    "strings"
)
```

Ejemplos de operaciones con cadenas de caracteres.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    fmt.Println(strings.Contains("Hello, World!", "Hello"))
    // Output: true

    r := strings.NewReplacer("<", "&lt;", ">", "&gt;")
    fmt.Println(r.Replace("This is <b>HTML</b>!"))
    // Output: This is &lt;b&gt;HTML&lt;/b&gt;!
}
```

Te invito a que visites la [documentación](https://pkg.go.dev/strings) de la liberia strings para que puedas ver todas las funciones que puedes usar.

## Paquetes

Con los paquetes podemos organizar más nuestro proyecto, un parque en Go sirve para contener nuestros archivos Go y organizar Para crear un paquete simplemente creas una carpeta, pero a si sola no indica nada para que sea un paquete tenemos que crear un archivo de Go y indicar dentro del archivo el nombre del paquete (nombre de la carpeta creada para paquete).

```go
📦Go_desde_0
 ┃ // Paquete (carpeta creada para paquete)
 ┣ 📂mensaje
 ┃ ┃ // Archivo dentro de paquete
 ┃ ┗ 📜saludo1.go
 ┃ // Archivo principal que trabaja con paquete main
 ┗ 📜main.go
```

Dentro de estos archivos tenemos que definir el paque y crear funciones. Para crear funciones podemos hacer de dos formas.

- **Privada:** En Go para crear un funcion o metodo privado simplemente inlciamos el nombre de la función con minúscula y esta función no puede ser accedida de fuera del paquete.
- **Pública:** En Go para crear un funcion o metodo publico solo inicias el nombre de la función con mayúscula V este funcione si puede ser accedida de fuera del paquete. Una función pública tiene que ser documentada Imciando_con su nombre.

Para poder continuar con esto, debemos de ejecutar el siguiente comando en nuestra terminal:

```sh
# Para inicializar un modulo de go
go mod init github.com/USER/NOMBRE_CARPETA_EN_LA_QUE_ESTAMOS_TRABAJANDO
```

Esto nos creara un archivo ``go.mod``

```go
// go.mod
module github.com/Wilovy09/Go_desde_0

go 1.22.1
```

Una vez creado esto, podemos empezar a crear la siguiente estructura de carpeta en nuestro proyecto.

```go
📦Go_desde_0
 ┃ // Creamos una carpeta "mensaje"
 ┣ 📂mensaje
 ┃ ┃ // Creamos un archivo "saludo1.go"
 ┃ ┗ 📜saludo1.go
 ┃ // Nuestro go.mod
 ┣ 📜go.mod
 ┃ // Nuestro main.go
 ┗ 📜main.go
```

En nuestro ``mensaje/saludo1.go`` crearemos 2 funciones, 1 privada y 1 pública.

```go
// mensaje/saludo1.go

// Definimos el paquete (nombre de la carpeta)
package mensaje

// Importamos lo que necesitamos
import "fmt"

// saludoPrivado es un mensaje privado
func saludoPrivado() {
    /*
     * Esto es una función privada, ya que el nombre de la función inicia con minusculas
     * El comentario que esta encima de la función es la "documentación" de la función
     * (lo que sale al hacer hover en VSCode)
    */
    fmt.Println("Hola desde mensaje/saludo1.go (saludoPrivado)")
}

// SaludoPublico es un mensaje público
func SaludoPublico() {
    /*
     * Esto es una función pública ya que el nombre de la función inicia con mayusculas
     * El comentario que esta encima de la función es la "documentación" de la función
     * (lo que sale al hacer hover en VSCode)
    */
    fmt.Println("Hola desde mensaje/saludo1.go (SaludoPublico)")
}
```

Ya con esto tenemos 2 funciones muy sencillas dentro de nuestro ``mensaje/saludo1.go``, ahora tenemos que mostrarlas en el ``main.go``.

El problema esta en que en ``main.go`` solo podemos acceder a funciones públicas, las funciones privadas solo se pueden acceder dentro del mismo paquete (no pueden salir de la carpeta donde se encuentra el archivo, las públicas si).

```go
// main.go
package main

import (
    "fmt"
    /*
     * Importamos el PAQUETE que creamos
     * github.com/USER/NOMBRE_CARPETA_EN_LA_QUE_ESTAMOS_TRABAJANDO/PAQUETE
    */
    "github.com/Wilovy09/Go_desde_0/mensaje"
)

func main() {
    fmt.Println("Hola desde main.go")

    // accedemos al paquete y usamos su metodo SaludoPublico
    mensaje.SaludoPublico()
    // mensaje.saludoPrivado() // Error: saludoPrivado is not visible
}
```

Ahora crearemos otro archivo en nuestro paquete ``mensaje``

```go
📦Go_desde_0
 ┣ 📂mensaje
 ┃ ┣ 📜saludo1.go
 ┃ ┃ // Creamos un nuevo archivo "saludo2.go"
 ┃ ┗ 📜saludo1.go
 ┣ 📜go.mod
 ┗ 📜main.go
```

Dentro de este nuevo archivo pondremos

```go
// mensaje/saludo2.go
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
```

Y modificaremos el ``main.go``

```go
// main.go
package main

import (
    "fmt"
    "github.com/Wilovy09/Go_desde_0/mensaje"
)

func main() {
    fmt.Println("Hola desde main.go")
    mensaje.SaludoPublico()
    // mensaje.saludoPrivado() // Error: saludoPrivado is not visible

    // llamamos a la nueva funcion de saludo2.go
    mensaje.OtraFuncion()
}
```

> [!WARNING]
> No podemos crear otra ``función main`` dentro de un paquete ya que la función main solo puede ser ``creada`` en el ``punto de entrada`` de nuestra aplicación. ( ``main.go`` )

## Maps

Las mapas son como una listas que almacena cantidad de datos con clave y valor, Map en Go se parece a los diccionarios de Python. Para definir una mapa se hace como variables 0 array, en este caso ponemos map luego dentro de los corchetes ponemos el tipo de dato que será la clave y luego el tipo de dato que será su valor y podemos hacer inicializando con valor predeterminado dentro de las llaves, clave y valor separados con dos puntos y cada elemento separados con coma.

```go
package main

import (
    "fmt"
)

func main() {

    // map[CLAVE]VALOR{ "CLAVE" : VALOR , "CLAVE" : VALOR , "CLAVE" : VALOR }
    // Definir un Map con un valor iniciado
    mapa1 := map[string]int{
        "uno":1,
        "dos":2,
        "tres":3,
    }
    fmt.Println(mapa1)

    // Para agregar datos al map
    // nombre[CLAVE] = VALOR
    mapa1["cuatro"] = 4
    fmt.Println(mapa1)

    // Buscar por clave
    fmt.Println(mapa1["cuatro"])

    // Eliminar un dato
    delete(mapa1, "cuatro")
    fmt.Println(mapa1)

    // Definir un map vacio
    mapa2 := make(map[int]string)
    mapa2[1] = "Uno"
    fmt.Println(mapa2)
}
```

## For y Range

El Range sirve para iterar sobre un array, slice, map, string, etc... y nos devuelve el indice y el valor de cada elemento.

```go
package main

import (
    "fmt"
)

func main() {

    array1 := []int{1,2,3,4}

    // para no iterar el indice ponemos _
    for _, num := range array1{
        fmt.Println(num)
    }

    // para iterar con el indice
    for i, num := range array1{
        fmt.Println(i, "=>", num)
    }
}
```

Aqui hay un ejemplo de como hacer un for con un map.

```go
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
```

## Type

En Go podemos crear un tipo de dato personalizado, para hacer esto usamos la palabra reservada ``type`` seguido del nombre del tipo de dato que queremos crear y el tipo de dato que va a ser.

```go
package main

import (
    "fmt"
)

type entero int

func main() {

    var a entero = 42

    fmt.Println(a)
}
```

## Struct

En Go podemos crear una estructura de datos personalizada, para hacer esto usamos la palabra reservada ``struct`` seguido del nombre de la estructura y dentro de las llaves ponemos los campos que va a tener nuestra estructura.

- Una estructura es un tipo de dato que puede contener varios campos.
- Los campos de una estructura pueden ser de cualquier tipo de dato.
- Para acceder a los campos de una estructura usamos el operador punto ``.``.
- La estructura la creamos fuera de la función main.

```go
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
```

## Métodos y herencia

Para poder seguir con esta parte, tenemos que tener el código desarrollado de la parte de [Struct](#struct).

```go
package main

import (
    "fmt"
)

func main() {
    curso1 := Curso{
        nombre: "Curso profesional de Go",
        url: "https://go.com/",
        habilidades: []string{"Go", "Backend"},
    }

}

type Curso struct {
    nombre string
    url string
    habilidades []string
}
```

### Crear métodos

Ahora aprenderemos a crear Métodos, los métodos se crean como una función pero con un receptor, el receptor es el nombre de la estructura que va a tener el método.

```go
package main

import (
    "fmt"
)

func main() {
    curso1 := Curso{
        nombre: "Curso profesional de Go",
        url: "https://go.com/",
        habilidades: []string{"Go", "Backend"},
    }

    // imprimimos la estructura curso1
    fmt.Println(curso1)

    // llamamos al método inscribirse
    curso1.inscribirse("Wilovy")

}
// Curso es una estructura que representa un curso
type Curso struct {
    nombre string
    url string
    habilidades []string
}

/*
 * inscribirse es un método de la estructura Curso
 * el receptor es el nombre de la estructura que va a tener el método
 * el método inscribirse recibe un parámetro nombre de tipo string
*/

// inscribirse es un método de la estructura Curso
func (c Curso) inscribirse(nombre string){
    fmt.Printf("La persona %s se ha inscrito al curso %s \n", nombre, c.nombre)
}
```

Si ponemos atención, estamos creando una función que se llama ``inscribirse`` y esta función tiene un receptor que es el nombre de la estructura que va a tener el método, en este caso es ``Curso``, entonces tenemos ``func (c Curso)`` que viene siendo básicamente una variable ``c`` de tipo ``Curso`` luego el nombre de la función y recibe como parametro un nombre de tipo string ``func (c Curso) inscribirse(nombre string){}``.

Y para acceder a los valores de la estructura usamos el receptor ``c``, en este ejemplo lo usamos para hacer un print del nombre del curso.

### Herencia

Ya sabemos crear métodos, ahora vamos a ver como hacer herencia en Go, para hacer herencia en Go simplemente creamos una estructura que tenga como campo otra estructura.

Básicamente vamos a heredar una estructura dentro de otra estructura.

```go
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
```

> [!TIP]
> Para entender de mejor forma la herencia en Go, te recomiendo que veas el video [Métodos y Herencia en Go - Roelcode](https://youtu.be/3dwTYkyR9_E?si=MwS4WsG474Q_tS8H).

---

## Estructura recomendada para proyectos en Go

Esta estructura es recomendada para proyectos en Go, para que sea más fácil de entender y de mantener.

Esta basada en la estructura recomendada en la documentaciónm oficial.

[Video de referencia](https://youtu.be/DHSp2VHP4dM?si=2SnVvGBU7bgCyiQJ)

```txt
📦estructura_recomendada
┣ 📂bin
┃ ┗ 📜myapp.exe
┣ 📂cmd
┃ ┣ 📂api
┃ ┃ ┗ 📜api.go
┃ ┗ 📜main.go
┣ 📂internal
┃ ┣ 📂config
┃ ┃ ┗ 📜env.go
┃ ┣ 📂db
┃ ┃ ┗ 📜db.go
┃ ┣ 📂middleware
┃ ┃ ┗ 📂auth
┃ ┃ ┃ ┣ 📜jwt.go
┃ ┃ ┃ ┗ 📜password.go
┃ ┣ 📂services
┃ ┃ ┗ 📂user
┃ ┃ ┃ ┣ 📜routes.go
┃ ┃ ┃ ┗ 📜store.go
┃ ┗ 📂types
┃ ┃ ┗ 📜types.go
┣ 📂pkg
┃ ┣ 📂middleware
┃ ┃ ┗ 📜cache.go
┃ ┗ 📂utils
┃ ┃ ┗ 📜ustils.go
┗ 📂test
  ┗ 📜test.go
```

- **bin:** Aquí se almacenan los binarios de nuestra aplicación.
- **cmd** Aplicación principal dentro de esta carpeta
  - No se pone demasiado código dentro de cmd.
  - Es común una función main pequeña.
  - Todo en paquetes.
- **internal** Código interno de la aplicación
  - Limita la visibilidad al propio paquete y subpaquetes.
  - Organiza el código interno sin exponer detalles de implementación.
  - Evita dependencias no deseadas y mantiene la separación de interfaces.
- **pkg** Código que puede ser importado por otros proyectos.
  - Almacena código reutilizable y compartible.
  - Mejora la claridad y organización del propyecto.
  - Facilita la separación entre la lógica de la aplicación y el código reutilizable.
- **test** Pruebas unitarias y de integración.
  - Pruebas unitarias y de integración.
  - Pruebas de rendimiento y de carga.
  - Pruebas de seguridad y de regresión.

## Testing con Go

Las pruebas en general suelen ser una parte importante de cualquier proyecto de software, ya que nos permiten asegurarnos de que nuestro código funciona correctamente.

### Prueba unitaria

Supongamos que tenemos una función que suma dos números y queremos probarla.

```go
// utils/utils.go
package utils

func Sum(a, b int) int {
    return a + b
}
```

Para probar esta función, creamos un archivo ``test/sum_test.go``.

```go
// test/sum_test.go
package test

import (
    "testing"
    "github.com/USER/Go_desde_0/utils"
)
func TestSum(t *testing.T) {
    result := utils.Sum(5, 5)
    expected := 10
    if result != expected {
      t.Errorf("Sum(5,5) expected %d but got %d", expected, result)
    } else {
      t.Logf("Sum(5,5) expected %d and got %d", expected, result)
    }
}
```

Ahora para ejecutar la o las pruebas tenemos que estar dentro de nuestra carpeta ``test``.

```sh
cd test
```

Una vez dentro de la carpeta ``test`` ejecutamos el siguiente comando.

```sh
go test
```

Si el test es exitoso nos mostrara algo asi:

```sh
# CONSOLE
$ go test
PASS
ok      github.com/Wilovy09/Go_desde_0/test     0.018s
```

De lo contrario:

```sh
# CONSOLE
$ go test
--- FAIL: TestSum (0.00s)
    sum_test.go:11: Sum(5,5) expected 11 but got 10
FAIL
exit status 1
FAIL    github.com/Wilovy09/Go_desde_0/test     0.027s
```

### Prueba de integración

Supongamos que tienes un servidor HTTP simple en Go que devuelve "Hola, mundo!" como respuesta:

```go
// main.go
package main

import (
    "net/http"
    "github.com/Wilovy09/Go_desde_0/utils"
)

func main() {
    http.HandleFunc("/", utils.Handler)
    http.ListenAndServe(":8080", nil)
}
```

```go
// utils/routes.go
package utils

import (
    "fmt"
    "net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "¡Hola, mundo!")
}
```

```go
// test/index_test.go
package test

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "github.com/Wilovy09/Go_desde_0/utils"
)

func TestHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(utils.Handler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    expected := "¡Hola, mundo!"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            rr.Body.String(), expected)
    }
}
```

Ahora para ejecutar la o las pruebas tenemos que estar dentro de nuestra carpeta ``test``.

```sh
cd test
```

Una vez dentro de la carpeta ``test`` ejecutamos el siguiente comando.

```sh
go test
```

Si el test es exitoso nos mostrara algo asi:

```sh
# CONSOLE
$ go test
PASS
ok      github.com/Wilovy09/Go_desde_0/test     0.022s
```

De lo contario:

```sh
# Console
$ go test
--- FAIL: TestHandler (0.00s)
    index_test.go:28: handler returned unexpected body: got ¡Hola, mundo! want ¡Hello, world!
FAIL
exit status 1
FAIL    github.com/Wilovy09/Go_desde_0/test     0.033s
```

## Go extras

```sh
# Para inicializar un modulo de go
go mod init github.com/USER/NOMBRE_CARPETA_EN_LA_QUE_ESTAMOS_TRABAJANDO
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
go doc NOMBRE_PAQUETE
```

```sh
# Para compilar un archivo de Go
go build NOMBRE_ARCHIVO.go
```

```sh
# Para ejecutar pruebas
go test
```
