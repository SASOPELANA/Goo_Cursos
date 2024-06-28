package main

import (
	"fmt"
	"runtime"
	"time" // Para saver la hora actual
)

func main() {

	fmt.Print("\n--------------------------------------------------------\n")
	fmt.Println("\nIf")
	t := time.Now() // Saca el tiempo actual
	hora := t.Hour() // Saca la hora actual
	
	if hora < 12 {
		fmt.Println("¡Mañana! Buenos días!")
	} else if hora < 19 {
		fmt.Println("¡Tarde! Buenas tardes!")
	} else{
		fmt.Println("¡Noche! Buenos noches!")
	}

	fmt.Print("\n--------------------------------------------------------\n")
	fmt.Println("\nSwitch")

	os := runtime.GOOS // Metodo para saver el sistema Operativo en Go

	switch os { // -> En Go no es necesario la instruccion break para cada caso
		case "windows":
			fmt.Println("Go run --> Windows")
		case "linux":
			fmt.Println("Go run --> Linux")
        case "darwin":
			fmt.Println("Go run --> MacOS")
		default:
			fmt.Println("Go run --> Otro OS")	
	}

	fmt.Print("\n---------------------------------------------------------------\n")
	fmt.Println("\nFunciones en Go.")
	hola()
	nombre("Sergio")
	saluda := devuelve("Alejandro")
	fmt.Println(saluda)

	fmt.Print("\n---------------------------------------------------------------\n")
	fmt.Print("\n *** Calculadora Basica ***\n")
	var a, b int

	fmt.Println("\nIngrese un número: ")
	fmt.Scanln(&a)
	fmt.Println("Ingrese otro número: ")
	fmt.Scanln(&b)

	suma, resta, producto, divi := calculadoraBasica(a, b)
	fmt.Println("La suma:", suma)
	fmt.Println("La resta:", resta)
	fmt.Println("Producto:", producto)
	fmt.Println("La division:", divi)

}


// Funcion que solo muestra 
func hola(){
	fmt.Println("Hola desde la función hola()")
}

// Funcion que pide un valor por parametro
func nombre(name string) {
	fmt.Println("Hola", name)
}

// Funcion que devuelve o retorna un valor
func devuelve(name string) string {
	return "Hola " + name
}

func calculadoraBasica (a, b int) (int, int, int, float32) {
	suma := a + b
	resta := a - b
	producto := a * b
	division := float32(a) / float32(b)
	return suma, resta, producto, division
}