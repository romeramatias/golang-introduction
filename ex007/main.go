package main

import "fmt"

// Calculo Las funciones tambien son tipos de datos
var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

// Funciones anonimas y closures
func main() {
	fmt.Println("5 + 7 es ", Calculo(5, 7))

	// La utilidad de definir una funcion de esta manera es que la puedo editar cuando quiera
	// como hacemos a continuacion
	Calculo = func(num1 int, num2 int) int {
		return num1 - num2
	}
	fmt.Println("6 - 4 es ", Calculo(6, 4))

	Calculo = func(num1 int, num2 int) int {
		return num1 / num2
	}
	fmt.Println("12 / 3 es ", Calculo(12, 3))

	Operaciones()

	// CLOSURES
	// Lo que hacemos aca es muy loco
	// tablaDel guarda un valor fijo en el que se van a multiplicar los nums de secuencia
	tablaDel := 2
	// miTabla guarda la funcion que retorna Tabla, osea, las lineas de codigo de la funcion Tabla
	// solo se ejecutan en esta asignacion, por ende numero siempre va a ser 2 y secuencia la arrancamos en 0
	miTabla := Tabla(tablaDel)

	for i := 1; i <= 10; i++{
		// Ahora que imprimimos lo que retorna la funcion interna de miTabla
		// Que va a ser una iteracion, porque cada vez que la llamamos, secuencia hace un +1 de su valor
		// numero siempre va a ser 2 y al incrementar secuencia, multiplica numero
		fmt.Println(miTabla())
	}

}

func Operaciones() {
	resultado := func() int {
		return 21 + 14
	}
	fmt.Println(resultado())
}

func Tabla(valor int) func() int{
	// Esta parte solo se ejecuta la primera vez que llamo
	numero := valor
	secuencia := 0

	// Y esta cada vez que tiene que retornar
	return func() int {
		secuencia += 1
		return numero * secuencia
	}

}
