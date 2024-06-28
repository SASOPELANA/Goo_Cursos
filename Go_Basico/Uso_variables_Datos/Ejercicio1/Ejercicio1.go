package main

import (
    "fmt"
    "math"
)

func main() {
	var lado1, lado2 float64

	// Entrada de datos en Go
	fmt.Print("Ingrese lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese lado 2: ")
	fmt.Scanln(&lado2)

	// Procesos
	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	// Salida de datos en Go
	fmt.Printf("\nÁrea: %.2f \n", area) // Muestra una salida de datos con dos decimales luego del punto   %.2f
	fmt.Printf("\nPerímetro: %.2f", perimetro) // Muestra una salida de datos con dos decimales luego del punto   %.2f
}