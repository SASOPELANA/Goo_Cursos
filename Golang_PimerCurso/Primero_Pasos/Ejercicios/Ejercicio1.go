// ? 1)
// * Crea 4 variables de diferentes tipos,
// * asignale a una tu edad,
// * a otra tu nombre,
// * a otra si te esta gustando el curso
// * y por ultimo el numero pi con 4 decimales

// ? 2)
// ! Imprime en pantalla el valor de las variables,
// ! con su tipo,
// ! y su representacion en binario, hexadecimal y octal

// * (Recomendaciones o tips)
// Usa el operador de declaracion corta
// Usa el tipo de dato adecuado para cada variable
// Usa un formato especifico en las impresiones

package main

import "fmt"

func Ejercicio1() {
	nombre := "Sergio"
	edad := 35
	curso := "Si me gusta el curso"
	numPi := 3.1415

	fmt.Println("Mi nombre: ", nombre)
	fmt.Println("Mi edad: ", edad)
	fmt.Println("Opinion del curso: ", curso)
	fmt.Println("Numero pi: ", numPi)

	fmt.Printf("\n%T, %T, %T, %T ", nombre, edad, curso, numPi)
	fmt.Printf("Valor binario edad: \n%b", edad)
	fmt.Printf("Valor Hexadecimal: \n%x, %x", numPi, edad)
	fmt.Printf("\nValor octal: %o ", edad)
}