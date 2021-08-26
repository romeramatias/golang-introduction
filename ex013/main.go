package main

import "fmt"

// Recursion funciones que se llaman a si misma
func main() {
	exponencia(2)
}

func exponencia(numero int) {
	// Con el return corto la ejecucion del metodo
	if numero > 100000000 {
		return
	}
	fmt.Println(numero)
	exponencia(numero * 2)
}
