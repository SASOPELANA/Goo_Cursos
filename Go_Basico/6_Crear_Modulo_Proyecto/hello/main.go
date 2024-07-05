package main

import (
	"fmt"
	"log"

	"github.com/SASOPELANA/Goo_Cursos/Go_Basico/6_Crear_Modulo_Proyecto/Mi_Modulo"
)

func main() {

	log.SetPrefix("Mi_Modulo: ")
	log.SetFlags(0)

	names := []string{"Sergio", "Juan", "Roberto"}
	messages, err := Mi_Modulo.Hellos(names)
	if err != nil {
		log.Fatal(err)
    }

	//message, err := Mi_Modulo.Hello("Sergio")
	//if err != nil {
	//	log.Fatal(err)
	//}
	fmt.Println(messages)

}