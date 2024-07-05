package main

import (
	"fmt"
	"log"

	"github.com/SASOPELANA/Goo_Cursos/Go_Basico/6_Crear_Modulo_Proyecto/Mi_Modulo"
)

func main() {

	log.SetPrefix("Mi_Modulo: ")
	log.SetFlags(0)

	message, err := Mi_Modulo.Hello("")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)

}