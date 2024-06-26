package main

import "fmt"

// Se considera una mala practica usar variables globales
var nombre2 string
var edad3 uint8
var peso float32
var bandera bool
var completo complex128

func main() {
	fmt.Println(nombre2)
	fmt.Printf("%T\n", nombre2)
	fmt.Println(edad3)
	fmt.Printf("%T\n", edad3)
	fmt.Println(peso)
	fmt.Printf("%T\n", peso)
	fmt.Println(bandera)
	fmt.Printf("%T\n", bandera)
	fmt.Println(completo)
	fmt.Printf("%T\n", completo)
}