package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	//funcionError()
	// Función para dividir dos numeros y manejo de errores
	//MostrarError()
	FuncionDefer()

}

// Muestra un error por que no se puede convertir una palabra a numero
func funcionError(){
	str := "123g" //  --> ERROR
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}else{
		fmt.Printf("El numero es: %d\n", num)
	}
	
}

// Función para encontrar un error en la division. Creando error.
func divide(divivendo int, divisor int) (int, error){
	if divisor == 0 {
		return 0, errors.New("No se puede dividir por cero.")
	}
	return divivendo / divisor, nil
}

func MostrarError(){
	// Función para dividir dos numeros y manejo de errores
	resultado, err := divide(10, 0) // --> NO SE PUEDE DIVIDIR CON CERO
	if err != nil {
		fmt.Println("Error:", err)
		return
	}else{
		fmt.Println("Resultado:", resultado)
	}
}

// Cierra una salida de flujo
func FuncionDefer(){
	file, err := os.Create("Hola.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	// TAMBIEN SIRVE PARA CERRAR UNA BASE DE DATOS.
	defer file.Close() // Se cierra el archivo al finalizar la función

	_, err = file.Write([]byte("Hola, Sergio Alejandro")) // Escribe en archivo txt creado
	if err != nil {
		fmt.Println(err)
		return
	}

}