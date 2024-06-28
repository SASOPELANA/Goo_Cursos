/*
	Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.
	El programa debe:
	Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.
	Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
	Calcular el área del triángulo usando la fórmula base x altura / 2.
	Calcular el perímetro del triángulo sumando los lados.
	Imprimir el área y el perímetro del triángulo con dos decimales de precisión.
	El programa debe usar variables para almacenar los lados del triángulo, la hipotenusa, el área y el perímetro. 
	También debe usar constantes para representar el número de decimales de precisión que se desean en la salida.
	Además, se deben utilizar funciones del paquete Math de Go para calcular la raíz cuadrada y cualquier otro cálculo matemático necesario.
*/

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