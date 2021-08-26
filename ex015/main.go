package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Al poner la instruccion 'go' siguie la ejecucion del programa, no se detiene a esperar
	// que miNombre termine, sigue con el main
	go miNombre("mati")
	fmt.Println("MATORRA")
	fmt.Println("MATORRAL")
	fmt.Println("MATO")
	fmt.Println("MATIENZO")
	fmt.Println("MATUTE")
	fmt.Println("MATUCHI")
	fmt.Println("MATEIKO")
	var x string
	fmt.Scanln(&x)
	// Hay que tener cuidado con el go porque no va a esperar que termine la ejecucion de ese metodo
	// Si el programa termina fuiste
}

func miNombre(nombre string) {
	letras := strings.Split(nombre, "")

	for _, letra := range letras{
		// Esto hace un set timeout de un segundo, cuando pasa el segundo seguimos con la siguiente instruccion
		fmt.Println(letra)
		time.Sleep(1000 * time.Millisecond)
	}
}
