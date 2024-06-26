package main 

import "fmt"

func main()  {
	// operador declaracion y asignacion en go      :=
	nombre := "Sergio"
	var apelleido string = "Sopelana"
	edad := 30
    fmt.Println(nombre, apelleido, edad)
	nombre = "Alejandro"
	fmt.Println(nombre)
}