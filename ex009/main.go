package main

import "fmt"

// Mapas
func main() {
	// Clave String, Valor String, se ve que se inicializa con el make
	paises := make(map[string]string) // make(map[string]string, 10) se le puede indicar los elementos

	paises["Mexico"] = "DF"
	paises["Argentina"] = "Buenos Aires"

	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	equiposPuntosMap := map[string]int{
		"River":     35,
		"PSG":       20,
		"Liverpool": 14,
	}
	// Lo ordena siempre alfabeticamente por la clave
	fmt.Println(equiposPuntosMap)

	// Agregar valores al mapa, busca la clave, si no la encuentra la crea
	equiposPuntosMap["Real Madrid"] = 22
	// Si ya existe en el mapa reemplaza el valor
	equiposPuntosMap["PSG"] = 30
	fmt.Println(equiposPuntosMap)

	// Para borrar elementos del mapa se usa el metodo delete
	delete(equiposPuntosMap, "Real Madrid")
	fmt.Println(equiposPuntosMap)

	// Aca el orden es distinto, el range no lo ordena alfabeticamente
	for equipo, puntaje := range equiposPuntosMap {
		fmt.Printf("El equipo %s tiene %d puntos \n", equipo, puntaje)
	}

	// Para preguntar si una clave existe en el mapa
	puntaje, existe := equiposPuntosMap["PSG"]
	fmt.Printf("El puntaje es %d, y el equipo existe '%t' \n", puntaje, existe)
}
