package libro

import "fmt"

// Polimorfismo en GO. Usando la palabra reservada interface, se debe teber algunos requisitos -->   PrintInfo()
type Printable interface {
	PrintInfo()
}

// Funcion para interface. Polimorfismo
func Print(p Printable){
	p.PrintInfo()
}



// Si colocamos las palabras el primer indice de la variable en mayuscula es un emcasulamiento Publico, si es en minuscula es Privado. Esto es en un struct
type Libro struct {
	titulo  string // ES PRIVADO CUANDO INICIA CON MINISCULA
	autor   string
	paginas int
}

// Simular Constructor en GO con puntero (Referencia)
// Con el constructor se accede a emcasulamiento privado.
func NuevoLibro(titulo string, autor string, paginas int) *Libro {
	return &Libro{
		titulo: titulo,
		autor:   autor,
        paginas: paginas,
	}

}

// Metodo para acceder a atrubutos privado en GO. Metodo get
func (b *Libro) SetTitulo(titulo string){
	b.titulo = titulo
}

// Metodo get para obtener. Para atributos privados en GO
func (b *Libro) GetTitulo()string{
	return b.titulo
}


func (b *Libro) PrintInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nPaginas: %d\n", b.titulo, b.autor, b.paginas)
}


// Composicion en GO. Una alternativa a la herencia en Go, ya que no tenemos el pilar de herencia en este lenguaje de programación.
type TextoLibro struct {
	Libro
	editorial string
	niveles string

}

// Constructor
func NuevoTexLibro(titulo, autor string, paginas int, editorial, niveles string) *TextoLibro{
	return &TextoLibro{
		Libro: Libro{titulo, autor, paginas},
		editorial: editorial,
		niveles: niveles,
	}
}

// Aplicando la alternativa a la herencia en Go, composicion
func (b *TextoLibro) PrintInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nPaginas: %d\nEditorial: %s\nNivel: %s\n", b.titulo, b.autor, b.paginas, b.editorial, b.niveles)
}