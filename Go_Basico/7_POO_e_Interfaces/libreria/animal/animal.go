// Interface en GO

package animal

import "fmt"

type Animal interface {
	Sonido()
}

// Estructura de perro y sus métodos
type Perro struct {
	Nombre string
}

func (p *Perro) Sonido(){
	fmt.Println(p.Nombre + "¡Hace guau guau!")
}

// Estructura de gato y sus métodos
type Gato struct{
	Nombre string
}

func (g *Gato) Sonido(){
	fmt.Println(g.Nombre + "¡Hace miau!")
}

// Función para imprimir el método sonido
func HacerSonido(animal Animal){
	animal.Sonido()
}