package main

import "fmt"

func fmtUno() {
	/*
	var edad = 35
	fmt.Println(edad)
	fmt.Printf("%T\n", edad) // Muestra el tipo de datos
	fmt.Printf("%b\n", edad) // Muestra el valor binario

	bandera := true
	fmt.Println("\n", bandera)
	fmt.Printf("Bool: %t\n", bandera) // Muestra el tipo de dato bool


	// Flotantes y complejos
	var flotante = 12.5
	var complejo = complex(12.5,15.5)
	fmt.Println(flotante)
	fmt.Println(complejo)
	*/

	// Diferencia entre print, printf, println.
	var imprimir = 25
	fmt.Println(imprimir)
	fmt.Print(imprimir)
	fmt.Printf("\n%T, %b, %v, %o, %x", imprimir, imprimir, imprimir, imprimir, imprimir)

}