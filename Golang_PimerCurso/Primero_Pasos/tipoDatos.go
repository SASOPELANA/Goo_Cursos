package main

import (
	"fmt"
)

// Variable globales, no es recomendable. Se considera una mala practica.
var nombre = "Juan"
var edad uint8 = 30


func Tipos_Datos() {
	fmt.Println(nombre)
	fmt.Printf("%T\n", nombre) //    "%T\n"   // Muestra el tipo de dato almacenado en al variable
	fmt.Println(edad)
	fmt.Printf("%T\n", edad) //    "%T\n"   // Muestra el tipo de dato almacenado en al variable

	// Go es lenguaje fuertemente tipado. 


	// Las comillas simples invertidas `` sirven para mostrar ese tipo de salida en la consola
	ciudadFavorita := `Mi
	Ciudad               Favorita              Oran          `
	fmt.Println(ciudadFavorita)
	fmt.Printf("%T\n", ciudadFavorita)
}