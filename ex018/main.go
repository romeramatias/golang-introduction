package main

import "fmt"

// Middlewares
// Interceptoores que permiten lehecutar instrucciones comunes a varias funciones que reciben y devuelven los mismos tipos de variables

var resultado int
func main() {
	fmt.Println("Inicio del programa")

	// A middleware le paso una funcion, se podra reutilizar con las funciones que reciben y devuelven los mismos
	// tipos de datos de las variables
	// A la derecha van los datos que le quiero enviar a la funcion que envie al middleware, porque este me la va a retornar
	resultado = middleware(sumar)(4, 13)
	fmt.Println(resultado)

	resultado = middleware(restar)(27, 6)
	fmt.Println(resultado)
}

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

// Recive la funcion con la cual voy a seguir la ejecucion del programa
// debo ponerle la firma con los tipos de datos que recive y devuelve
// y la devolucion del middleware sera esta misma funcion, la firma lo mismo
func middleware(f func(int, int) int) func(int, int) int {
	// Retorno una funcion anonima que es la funcion que recive los paremetros de la que me llega por paremetros
	return func(a, b int) int {
		// Ejecuto el middleware
		fmt.Println("Middleware")
		// Y retorno la funcion sumar/restar
		return f(a, b)
	}
}
