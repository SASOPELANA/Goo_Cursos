package main

import (
    "fmt"
     "github.com/SASOPELANA/Goo_Cursos/Go_Basico/6_Crear_Modulo_Proyecto/Mi_Modulo"
)

func main() {

	message := Mi_Modulo.Hello("Sergio")
	fmt.Println(message)

}