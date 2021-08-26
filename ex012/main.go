package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	readFile2()
	writeFile()
	writeFile2()
}

func readFile() {
	// Le tuve que dar el path absoluto porque no se la razon del error
	file, err := ioutil.ReadFile("C:\\Chattigo\\Workspace\\CursoGo\\ex012\\files\\archivo.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(file))
	}
}

func readFile2() {
	// Con el paquete os no solo puedo leer el archivo sino que tambien puedo editarlo
	file, err := os.Open("C:\\Chattigo\\Workspace\\CursoGo\\ex012\\files\\archivo.txt")

	if err != nil {
		fmt.Println(err)
	} else {
		scanner := bufio.NewScanner(file)
		// Esto lee por lineas
		for scanner.Scan() {
			reg := scanner.Text()
			fmt.Printf("Linea > " + reg + "\n")
		}
	}
	// Siempre hay que cerrarlo
	// Siempre hay que cerrarlo
	file.Close()
}

func writeFile() {
	// Crea un nuevo archivo
	file, err := os.Create("C:\\Chattigo\\Workspace\\CursoGo\\ex012\\files\\archivo2.txt")
	if err != nil {
		fmt.Println(err)
		// El return vacio corta la ejecucion a lo break
		return
	} else {
		// Con esta funcion se graba en un archivo, borra el contenido que ya existe
		fmt.Fprintln(file, "Linea nueva")
		file.Close()
	}
}

func writeFile2(){
	path := "C:\\Chattigo\\Workspace\\CursoGo\\ex012\\files\\archivo2.txt"
	if appendLines(path, "\nNueva linea") == false {
		fmt.Println("Error al escribir linea")
	}
}

func appendLines(path string, texto string) bool {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = file.WriteString(texto)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
