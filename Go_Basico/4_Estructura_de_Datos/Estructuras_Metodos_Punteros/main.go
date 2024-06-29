package main

import "fmt"


func main() {
	MostrarPersona()
	
	// Puntero
	var x int = 10
	fmt.Println(x)

	editar(&x)
	fmt.Println(x)

	// Crear una instancia persona y llamar a un metodo
	p1 := Persona{"Daniel", 20, "josechad@gmail.com",}
	p1.saluda() // -->    Metodo saludar

}

// Funcion con puntero
func editar(x *int){
	*x = 20
}



// Se crea una estructura persona. Similiar a crear una clase Persona en otro lenguaje
type Persona struct {
	nombre string
	edad int
	correo string
}

func MostrarPersona(){
	// Crear una instancia de la estructura Persona
	var p  Persona
	p.nombre = "Sergio"
	p.edad = 35
	p.correo = "sopeko@gmail.com"
	fmt.Println(p)

	// Crear otra instancia de la misma estructura Persona
	p2 := Persona{"Ajenadro", 25, "alejandro2024@gmail.com",}
	fmt.Println(p2)
}


// Crear un metodo
func (p3 *Persona) saluda(){
	fmt.Println("Hola, mi nombre es", p3.nombre)
}