package main

import (
	"fmt"
	"libreria/libro"
)

func main() {
	/*
	var MiLibro = libro.Libro{
		Titulo: "Programación en GO",
		Autor: "Google",
		Paginas: 300,
	}
	*/	

	// MiLibro.PrintInfo()

	// Usando Constructor
	fmt.Printf("\nUsando Constructor\n")
	MiLibro1 := libro.NuevoLibro("Mody Dick", "Herman Melville", 300)
	MiLibro1.PrintInfo()

	// Acceder a lo atributos privados
	fmt.Printf("\nUsando Metodo Set y Get, para acceder a un atributo privado.\n")
	MiLibro1.SetTitulo("Moby Dick (Edición especial)")
	fmt.Println(MiLibro1.GetTitulo())

	MiLibro1.PrintInfo()
}