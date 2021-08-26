package main

import (
 "fmt"
	"strconv"
)

// Global scoping para todas las funciones
var numero int

// Nombre Si la primer letra es en miniscula no se exporta, solo se puede usar aca
// Lo mismo con las funciones, hay que poner la primer letra en mayuscula
var Nombre string = "Matias"
var status bool = true

func main() {
	// Te inicializa las variables al declararlas
	fmt.Println(numero)
	fmt.Println(Nombre)
	fmt.Println(status)

	// Scope local de la funcion
	var numero2 int
	fmt.Println(numero2)

	// Al inicializarla evitas indicar el tipo de dato y el var
	numero3 := 3
	numero3 = 4
	fmt.Println(numero3)

	// Declaracion de variables encadenada
	var num1, num2, num3 int
	// Inicializacion encadenada de variables
	// Variable 'status' local, mismo nombre que la global, pero nueva al poner :=
	asd1, asd2, status := 1, 2, false
	fmt.Println(num1, num2, num3)
	fmt.Println(asd1, asd2, status)

	// Llamada a una funcion
	MostrarNombre()

	// -- CASTEOS --
	var numerito float32 = 44
	// Para el casteo se utiliza el tipo de variable como una funcion
	num1 = int(numerito)
	// Paquete que castea los nums en strings
	fmt.Println(strconv.Itoa(num1))
}

// MostrarNombre se exporta al estar en mayusculas
func MostrarNombre(){
	fmt.Println(Nombre)
}