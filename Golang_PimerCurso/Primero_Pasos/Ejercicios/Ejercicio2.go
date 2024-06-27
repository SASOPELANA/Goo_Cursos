// * Declara 3 variables, una de tipo entero, otra de tipo cadena y otra de tipo booleano.
// * Con los identificadores de edad, nombre y si ya dejaste tu resena de 5 estrellas


// ! Puntos a tomar en cuenta
// Las variables tienen que ser variables globales
// No puedes usar el operador de declaracion corta
// No puedes declarar e iniciarlizar las variables en la misma linea
// Imprime el valor asignado a cada variable sin asignar ningun valor, asigna los valores y vuelve a imprimirlos
// como se llama el valor que se le asigna a las variables al ser declaradas (sin ser inicializadas)

package main

import "fmt"

// Esto es una mala practica, solo se debe aplicar en un caso espesifico. Declarar variables globales.
var edad int
var nombre string
var bandera bool

func main() {

	fmt.Println(edad, nombre, bandera)

	fmt.Printf("\n%v, %v, %v\n", edad, nombre, bandera)
	edad = 25
	nombre = "Sergio"
	bandera = true
	fmt.Printf("\n%v, %v, %v\n", edad, nombre, bandera)
}