package main

import (
	"fmt"
	"libreria/animal"
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

	// Usando Composicion en Go. Una alternativa a la Herencia en este leguanje de programación.
	fmt.Printf("\nUsando Composicion en GO. Alternativa a la Herencia\n")

	MiTextoLibro := libro.NuevoTexLibro("Comunicación", "Jaime Gamarra", 250, "Santillana SAC", "Terciario")
	MiTextoLibro.PrintInfo()

	// Polimorfismo en GO
	fmt.Printf("\nPolimorfismo en GO\n")

	//MiLibro1.PrintInfo()
	//MiTextoLibro.PrintInfo()

	libro.Print(MiLibro1)
	fmt.Print("\n")
	libro.Print(MiTextoLibro)

	// Haciendo uso de Interfaces en GO
	fmt.Print("\n-----------------------------------------------------------\n")
	fmt.Printf("\nInterfaces en GO\n")

	miPerro := animal.Perro{Nombre: "Firulais "}
	miGato := animal.Gato{Nombre: "Tom "}

	// Enviar por puntero o referencia.
	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	// Iterar la lista o vector de animales
	fmt.Print("\n-----------------------------------------------------------\n")
	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max "},
		&animal.Gato{Nombre: "Sam "},
		&animal.Perro{Nombre: "Cazador "},
		&animal.Gato{Nombre: "Silvestre "},
	}

	for _, animal := range animales{
		animal.Sonido()
	}

}
