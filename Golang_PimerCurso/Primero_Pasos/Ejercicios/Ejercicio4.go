// * Crea tu propio tipo de dato que provenga de un entero para la edad (Lo mas optimo posible)
// * Crea una variable con tu nuevo tipo con el identificador "edad" usando sin el operador de declaracion corta
// * Imprime el valor de la variable "edad", el tipo y asignale el valor de 20
// * imprime el valor de edad.

package main

import "fmt"

func main() {
	type edad1 uint8
	var edad  edad1

	fmt.Printf("\n%v, %T", edad, edad)
	edad = 35
	fmt.Printf("\n%v", edad)

	fmt.Println(" ")
	fmt.Print("\n------------------------\n")
	fmt.Println("Ejercicio 5 ")
	Ejercicio5()
}

// * Del ejercicio anterior. Crea una variable global sin el operador de asignacion corta, con el identificador de
// * EdadPersona con un tipo de dato que provenga de Edad.
// * Usa la conversion para pasar el valor guardado en edad con el tipo de dato que proviene.
// * Utiliza el operador de declaracion corta para asignarle el valor a EdadPersona.
// ! Copia y pega el codigo del ejercicio 4



func Ejercicio5() {

	type Edad uint8
	var edad Edad
	var edadPersona uint8

	fmt.Printf("%v\n",edad)
	fmt.Printf("%T\n",edad)
	edad = 20
	fmt.Printf("%v\n",edad)

	fmt.Printf("%v, %T\n",edadPersona,edadPersona)
	edadPersona = uint8(edad)
	fmt.Printf("%v, %T\n",edadPersona,edadPersona)
	edadPersona = 10
	fmt.Printf("%v, %T\n",edadPersona,edadPersona)

}

