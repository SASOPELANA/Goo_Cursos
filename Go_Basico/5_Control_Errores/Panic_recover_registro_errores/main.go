package main

import (
	"fmt"
	"log" // Para mendajes de registros
	"os"
)

func main() {
	//divide(100, 10)
	//divide(600, 30)
	//divide(300, 0)
	//divide(500, 200)


	// Mensajes de registros. con funcion  -->    log()  registra errores
	log.Print("Este en un medaje de registro.")
	log.Println("Este es otro mensaje de registro.")

	// Crea un archivo y guarda los errores
	log.SetPrefix("main():") // muestra en que funcion se genero el registro
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)

	if err != nil{
		log.Fatal(err)
	}

	defer file.Close() // Cierra el archivo

	log.SetOutput(file)
	log.Print("!Oye, soy un Log")

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