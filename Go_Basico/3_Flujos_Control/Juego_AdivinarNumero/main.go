package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"github.com/fatih/color"
)

func main() {
	// rand.Intn()  --> Funcion randon en Go

	verMenu()

}

func Jugar() {
	numAleatorio := rand.Intn(100) + 1
	var numIngresado int
	var intentos int
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		color.Cyan("Ingrese un número (intentos restantes: %d): ", maxIntentos - intentos + 1)
		fmt.Scan(&numIngresado)

		if numIngresado == numAleatorio {
			color.Red("Felicitaciones! Has ganado! Advinaste el número!")
			JugarNuevamente()
            return // Para que salga del for
		} else if numIngresado < numAleatorio {
			color.Green("El número a adivinar es mayor.")
		} else if numIngresado > numAleatorio {
			color.Yellow("El número a adivinar es menor.")
		}
	}

	fmt.Println("Perdiste! El número era:", numAleatorio)
	JugarNuevamente()
}

func JugarNuevamente(){
	var eleccion string
	color.Magenta("¿Desea jugar? (S/N): ")
	fmt.Scan(&eleccion)

	if eleccion == "S" || eleccion == "s" {
		limpiarConsola()
        Jugar()
    } else if eleccion == "N" || eleccion == "n" {
        color.Magenta("Gracias por jugar!")
    } else{
		fmt.Println("Opción inválida. Intente nuevamente.")
		JugarNuevamente()
	}

}

func verMenu() {
	limpiarConsola()
	color.Blue("\n----------------------------------------------------------------\n")
	color.Cyan("\n                 ********   MENÚ  *********")
	color.Blue("\n----------------------------------------------------------------\n")

	color.Red("\n        ♥   ♥   ♥  --- REGLAS DEL JUEGO ---  ♥   ♥   ♥     ")
	color.Green("\n↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔")
	color.Yellow("\n El juego consite en adivinar un número aleatorio del 1 al 100.")
	color.Yellow(" Tiene hasta diez intentos. De no adivinar, pierde el juego.")
	color.Green("\n↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔↔")
	fmt.Println("\n ")

	color.Magenta("1. Jugar")
	color.Green("2. Creditos")

	var eleccion int

	color.Yellow("Ingrese su elección (1/2): ")
	fmt.Scan(&eleccion)
	switch eleccion {
		case 1:
			limpiarConsola()
			Jugar() // Llama a la función que simula el juego del adivinar número 1 al 100

		case 2:
			limpiarConsola()
            Creditos() // Llama a la función que muestra los créditos del programa	
		default:
			fmt.Println("Opción inválida. Intente nuevamente.")
            verMenu() // Llama a la función que muestra el menú de nuevo
	}

	// fmt.Println("Presiona Enter para continuar...")
    // fmt.Scanln()

}

func limpiarConsola() {
    // Verificar el sistema operativo para determinar el comando adecuado
    switch runtime.GOOS {
    case "windows":
        clearCmd := exec.Command("cmd", "/c", "cls") // Comando para limpiar en Windows
        clearCmd.Stdout = os.Stdout
        clearCmd.Run()
    case "linux", "darwin":
        clearCmd := exec.Command("clear") // Comando para limpiar en Linux y macOS
        clearCmd.Stdout = os.Stdout
        clearCmd.Run()
    default:
        fmt.Println("No se pudo limpiar la consola. Sistema operativo no soportado.")
    }
}

func Creditos(){
	limpiarConsola()
	hiCyan := color.New(color.FgHiCyan)
	color.Green("\n°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n ")
	hiCyan.Println("   Programador: Sergio Alejandro Sopelana")
	color.Green("\n°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°°\n ")
	JugarNuevamente()
}