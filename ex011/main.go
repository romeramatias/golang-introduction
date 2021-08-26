package main2

import "fmt"

// Interfaces

func main() {
	Matias := new(hombre)
	humanoRespirando(Matias)

	Jose := new(mujer)
	humanoRespirando(Jose)
}

// Dentro de la interface se definen los metodos
type humano interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal interface {
	ClasificacionVegetal() string
}

// Estructuras
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}
type mujer struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
}

// Implementacion de los metodos de humano en la estructura Hombre
// Al parecer no hace falta indicar que implementa la interface
// No se si debe implementar todos los metodos o solo los que escribo
func (h *hombre) respirar()    { h.respirando = true }
func (h *hombre) comer()       { h.comiendo = true }
func (h *hombre) pensar()      { h.pensando = true }
func (h *hombre) sexo() string { return "Hombre" }

func (m *mujer) respirar()    { m.respirando = true }
func (m *mujer) comer()       { m.comiendo = true }
func (m *mujer) pensar()      { m.pensando = true }
func (m *mujer) sexo() string { return "Mujer" }

// Polimorfismo, no importa si le llega un hombre o una mujer
func humanoRespirando(hu humano) {
	hu.respirar()
	fmt.Println("Soy un", hu.sexo(), "y ya estoy respirando")
}
