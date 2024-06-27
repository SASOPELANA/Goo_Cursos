package main

import (
	"fmt"
	"rsc.io/quote"
)

// Declaracion de constantes a nivel global. Esto se considera una mala practica.
// Con las variables constantes si se puede ejecutar la aplicacion sin hacer uso de la misma. Se debe guardar con mayuslas las constantes.
const Pi = 3.14

// Declaracion de varias constantes
const (
	x = 100 // Decimal
	y = 0b1010 // Binario
	z = 0o12   // Octal
	w = 0xFF   // Hexadecimal
)

// Uso de la funcion iota para asignar numeros correlativos automaticamnete. Ejemplo
const (
	Domingo = iota + 1 // iota incia de cero (0), por defecto.
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)


func main() {
	fmt.Println("Hola Mundo!")
	fmt.Println(quote.Go())

	// Declaración de variables
	// var firstName, lastName string
	// var age int
	// var firstName, lastName, age = "Sergio", "Alejandro", 35       // Tambien se puede declarar asi.
	// firstName, lastName, age := "Sergio", "Alejandro", 35       // Tambien se puede declarar asi.  Operador de declaracion y asignacion en Go   :=

	var ( // Otra forma de declarar variables en Go
		firstName, lastName string
		age int
	)

	firstName = "Sergio"
	lastName = "Alejandro"
	age = 35
	fmt.Println(firstName, lastName, age)

	fmt.Println(Pi)
	fmt.Println(x, y, z, w)

	fmt.Println(Jueves)

}