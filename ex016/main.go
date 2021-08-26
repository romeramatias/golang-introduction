package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Los canales se declaran con el make(chan y un tipo de dato) que tiene que ser del tipo que va a guardar luego
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("Estamos acá")

	// El await de go, sin esto la ejecucion termina, al enviarle un canal y con esta instruccion
	// tenemos que esperar que devuelva ese dato
	// canal se llena del dato que esta en a funcion bucle
	//tiempoTranscurrido := <-canal1
	//log.Println(tiempoTranscurrido)
	log.Println(<-canal1)
}

// bucle le enviamos por parametros el canal que creamos en el main para guardar una informacion que esta funcion va a tener
func bucle(canal chan time.Duration) {
	horarioInicioBucle := time.Now()

	// Iteracion para que solamente pase el tiempo
	for i := 0; i < 100000000000; i++ {

	}

	// Guardamos tambien en una variable la hora en la que termina el bucle for
	horarioFinBucle := time.Now()

	// Guardamos el tiempo que tarda en ejecutarse el bucle
	canal <- horarioFinBucle.Sub(horarioInicioBucle)

}
