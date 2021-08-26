// Package persona: Mismo nombre que el ombre de la clase
package persona

import "time"

// Persona Con mayusucaslas asi se puede exportar, por el scoping
type Persona struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Status    bool
}

func (this *Persona) NewPersona(id int, nombre string, fechaAlta time.Time, status bool) {
	// Como un constructor, el this indica el puntero hacia la persona a la cual se seteamos los valores
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaAlta
	this.Status = status
}
