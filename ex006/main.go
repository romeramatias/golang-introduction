package main

import "fmt"

func main() {
	fmt.Println(primerFuncion(21))

	// Para guardar ambas devolciones
	numero, estado := segundaFuncion(2)
	fmt.Println(numero, estado)

	fmt.Println(tercerFuncion(3))
	fmt.Println(cuartaFuncion(3, 4, 6))
}

// Firma de funciones en go
func primerFuncion(numero int) int {
	return numero * 2
}

// Devolucion de multiples valores de distintos tipos de datos
func segundaFuncion(numero int) (int, bool) {
	if numero == 1 {
		return 5, true
	} else {
		return 10, false
	}
}

// Declaro el nombre de la variable que voy a retornar
func tercerFuncion(numero int) (retorno int){
	retorno = numero * 2
	return
}

// Parametros variables, con los puntos indicamos que no sabemos cuantos parametros llegaran
// Go no tiene sobrecarga de metodos
func cuartaFuncion(numeros ...int) int {
	total := 0
	// Inicializo dos variables que se llenan de los numeros que envian por parametros
	// Con el _ es como que descartas la variable, como en Kotlin
	for _, num := range numeros {
		// El range devuelve dos cosas
		// En este ejemplo el segundo return del range es el numero que enviamos por parametros
		// Aca hago una sumatoria de los numeros que me enviaron
		total += num
	}
	return total
}