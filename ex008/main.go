package main

import "fmt"

// Arreglos
func main() {
	// Ya se inicilizan las diez posiciones en cero, esta bueno por los null pointer
	var array [10]int
	array[0] = 14
	array[1] = 21
	array[2] = 13
	fmt.Println(array)

	array2 := [5]int{2, 6}
	fmt.Println(array2)

	// El len es una funcion a diferencia de otros lenguajes que era una propiedad
	for i := 0; i < len(array2); i++ {
		fmt.Println(array2[i])
	}

	// MATRICES
	var matriz [5][7]int
	matriz[1][2] = 5
	fmt.Println(matriz)

	// SLICES
	// Son vectore dinamicos, a lo array list
	// Si le ponemos la longitud a los corchetes es un vector/array
	slice := []int{21, 14, 15, 13, 17, 22, 27}
	sliceCortado := slice[2:]
	fmt.Println(slice)
	fmt.Println(sliceCortado)
	// Funcion make para crear slices
	// Se le envia, el tipo, el largo y la capacidad reservada por si lo queremos ampliar
	sliceTercero := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d", len(sliceTercero), cap(sliceTercero))

	sliceCuarto := make([]int, 0, 0)
	for i := 0; i < 20; i++ {
		sliceCuarto = append(sliceCuarto, i)
	}
	// Al no tener la capacida definida va reservando espacio en memoria automaticamente de manera binaria
	fmt.Printf("\nLargo %d, Capacidad %d", len(sliceCuarto), cap(sliceCuarto))
}
