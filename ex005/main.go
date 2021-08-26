package main

import "fmt"

// No existe el while y el do while, toda iteracion se maneja con el for

func main() {
	i := 1

	// La condicion al parecer hay que ponerla dentro
	for i < 10 {
		fmt.Println(i)
		i++
	}

	// For clasico
	for i := 0; i <= 20; i++ {
		fmt.Println(i)
	}

	// For de loop indefinido con break
	numero := 0
	for {
		fmt.Println("Sigo")
		fmt.Scanln(&numero)

		// Si ingresa el numero 100 rompemos la ejecucion
		if numero == 100 {
			break
		}
	}

	// For con continue
	var j = 0
	for j < 10 {
		fmt.Printf("\n Valor de j: %d", j)
		if j == 5 {
			fmt.Printf(" el numero es cinco, lo multiplicamos por 2")
			j *= 2
			continue
		}
		fmt.Printf(", Pasamos por aca")
		j++
	}

	// Goto labels
	x := 0

RUTINA:
	for x < 10 {
		fmt.Println("X =", x)
		if x == 4 {
			x = x + 2
			fmt.Println("Vuelvo")
			goto RUTINA
		}
		x++
	}

}
