package main

import "fmt"

func vaar() {	
	nombre := "Sergio"
	edad := 25
	fmt.Println(nombre, edad)

	var apellido = "Sanchez"
	var comidaRapida string
	fmt.Println(nombre, apellido, edad)
	comidaRapida = "Mila napolitana"
	fmt.Println(nombre, apellido, edad, comidaRapida)

	fmt.Println(" ")
	nombre1()
}	


func nombre1()  {
	// operador declaracion y asignacion en go      :=
	nombre := "Sergio"
	var apelleido string = "Sopelana"
	edad := 30
    fmt.Println(nombre, apelleido, edad)
	nombre = "Alejandro"
	fmt.Println(nombre)
}