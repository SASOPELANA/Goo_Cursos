package main

import "fmt"

func main() {
	divide(100, 10)
	divide(600, 30)
	divide(300, 0)
	divide(500, 200)

}

// Uso de funcion panic, para generar un pila de errores, por entrar en panico la aplicación
func divide(dividendo int, divisor int) {

	// Funcion anonima para capturar el panico. Para no interrupir con la ejecucion del programa.
	defer func ()  {
		if r := recover(); r != nil{
			fmt.Println(r)
		}
	}() // Fin de la funcion anonima

	validarCero(divisor)
	fmt.Println(dividendo / divisor)
}

// Funcion panic
func validarCero(divisor int) {
	if divisor == 0 {
		panic("No puedes dividir por cero.")
	}
}