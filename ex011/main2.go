package main2

import "fmt"

// Interfaces

func main2() {
	Matias := new(hombre2)
	Matias.esHombre = true
	humanoRespirando(Matias)

	Jose := new(mujer2)
	humanoRespirando(Jose)
}

// Dentro de la interface se definen los metodos
type humano2 interface {
	respirar()
	pensar()
	comer()
	sexo() string
}

type animal2 interface {
	respirar()
	comer()
	EsCarnivoro() bool
}

type vegetal2 interface {
	ClasificacionVegetal() string
}

// Estructuras
type hombre2 struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	esHombre   bool
}
type mujer2 struct {
	hombre2
}

// Implementacion de los metodos de humano en la estructura Hombre
// Al parecer no hace falta indicar que implementa la interface
// No se si debe implementar todos los metodos o solo los que escribo
func (h *hombre2) respirar() { h.respirando = true }
func (h *hombre2) comer()    { h.comiendo = true }
func (h *hombre2) pensar()   { h.pensando = true }
func (h *hombre2) sexo() string {
	if h.esHombre {
		return "Hombre"
	} else {
		return "Mujer"
	}
}

func (m *mujer2) respirar()    { m.respirando = true }
func (m *mujer2) comer()       { m.comiendo = true }
func (m *mujer2) pensar()      { m.pensando = true }
func (m *mujer2) sexo() string { return "Mujer" }

// Polimorfismo, no importa si le llega un hombre o una mujer
func humanoRespirando2(hu humano) {
	hu.respirar()
	fmt.Println("Soy un", hu.sexo(), "y ya estoy respirando")
}
