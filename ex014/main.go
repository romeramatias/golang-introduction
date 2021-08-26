package main

import (
	"fmt"
	"log"
	"os"
)
// Manejo de errores
func main() {
	// exampleDefer()
	examplePanic()
}

func exampleDefer(){
	// El defer es una instruccion que se va a ejecutar si o si cuando una funcion finaliza por return, error, o por lo que sea
	// No importa el orden, no va secuencialmente, se ejecuta al final siempre
	file := "archivo.txt"
	f, err := os.Open(file)

	// -- Esto por ejemplo no se ejecuta
	// -- defer fmt.Println("Surgio un error")
	// El close libera el bufer en memoria
	defer f.Close()


	if err != nil {
		fmt.Println("Errorazo")
		os.Exit(1)
	}
}

func examplePanic() {
	// No andaba el fmt porque el defer ejecuta solo una cosa
	defer func() {
		// Lo que hace el recover es traer el resultado del panic, como un catch
		reco := recover()

		// Si reco no es null es porque hubo un panic
		if reco != nil {
			log.Fatalf("Error %v", reco)
		}
	}()

	i := 10

	// Es como un runtime exception, se controla con el recover
	if i == 10 {
		panic("i VALE 10!")
	}
}
