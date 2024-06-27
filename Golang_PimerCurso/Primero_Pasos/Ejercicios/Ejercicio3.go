// * Del ejercicio anterior asigna los valores de 50, "Capdesis", true (al tipo de dato que corresponda)

// ! Usa el String print (recuerda del paquete fmt), imprime todos los valores con un solo print *
// ! y en la misma linea (agregale un formato).
// ! Asignaselos a una nueva variable que se llame resultado.

package main

import "fmt"

// Esto es una mala practica, solo se debe aplicar en un caso espesifico. Declarar variables globales.
var edad int
var nombre string
var bandera bool

func ejercicio3() {

	fmt.Println(edad, nombre, bandera)

	edad = 50
	nombre = "Capdesis"
	bandera = true
	
	resultado := fmt.Sprintf("\t%v, \t%v, \t%v\n", edad, nombre, bandera)
	fmt.Println(resultado)
}