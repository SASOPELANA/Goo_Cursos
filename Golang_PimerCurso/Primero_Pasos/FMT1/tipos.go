package main

import "fmt"

func main() {
	var x int
	type edad uint8
	var y edad
	x = 20
	y = 33
	fmt.Printf("Variables Tipo: %T, Valor: %v\n", x, x)
	fmt.Printf("Variables Tipo: %T, Valor: %v\n", y, y)
	
	// x = y
	// No se puede hacer, por que son de distintos tipos. Go es altamente tipado y dinamico.


	// * Esto si se puede hacer en Go. Conversiones
	x = int(y)
	fmt.Println("\nNuevo valor de x: ", x)
	y = edad(x)
	fmt.Println("\nNuevo valor de y: ", y)

	// Ejemplo de conversiones
	var entero int = 10
	var flotante float64 = 50.5
	var cadena string = "hola"

	fmt.Println("\n", entero, flotante, cadena)
	entero = int(flotante)
	fmt.Println("\n", entero)
	entero = 25
	flotante = float64(entero)
	fmt.Println("\n", flotante)
}

