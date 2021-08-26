package main

import "fmt"

var estado = true

func main() {
	estado = false

	// No necesita la condicion en parentesis
	if estado == false {
		fmt.Println("Estado es false")
	} else {
		fmt.Println("Estado es true")
	}

	// Se puede asignar un valor a una variable en la condicion
	if estado = true; estado == true {
		fmt.Println("Estado es true")
	}

	// SWITCH
	// No hace falta el break
	// Al declarar la variable asi, solo sirve para el switch
	switch numero := 4; numero {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("Otro num")
	}
}
