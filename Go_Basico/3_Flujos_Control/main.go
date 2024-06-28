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
}


func hola(){
	fmt.Println("Hola desde la función hola()")
}