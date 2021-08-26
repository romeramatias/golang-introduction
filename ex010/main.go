package main

import (
	"fmt"
	// Para poder importar el paquete (carpetas) tuve que crear el go.mod, los ordena alfabeticamente
	per "go-multitenancy/persona"
	"time"
)

// En go no existe la oop, en go se basa en las estructuras

func main() {
	// Matias := new(persona)
	// Matias.Id = 21
	// Matias.Nombre = "Matias"

	// Las propiedades que no les pongo valor igualmente se inicializan, no hay nulos
	// fmt.Println(Matias)

	atleta1 := new(atleta)
	atleta1.NewPersona(1, "Usain", time.Now(), false)
	fmt.Println(atleta1.Persona)
}

// Se puede aplicar herencia
type atleta struct {
	// Tiene como un puntero a Persona
	// Se pueden agregar propiedades
	per.Persona
}
