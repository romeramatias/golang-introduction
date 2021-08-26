package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var leyenda string

func main() {

	fmt.Println("Ingrese un numero:")
	fmt.Scanln(&numero1)

	fmt.Println("Ingrese otro numero:")
	fmt.Scanf("%d", &numero2)

	// BUFIO
	fmt.Println("Ingrese su nombre:")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()
	}

	fmt.Println(leyenda, ", la suma de ambos numeros es:", numero1+numero2)

}
