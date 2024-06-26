package main

import "fmt"

// En carpeta solo puede tenEr un solo archivo main pricipal en GO
func HolaMUndo()  {
	fmt.Println("Hola Mundo en GO")
	fmt.Println(" ")
	fmt.Println(8+10)
	fmt.Println(15/3)
	fmt.Println(!true)
	fmt.Println(false || true)
	fmt.Println(" ")
	adios()
}

// Funcion en GO
func adios() {
	fmt.Println("Adios Mundo")
}
