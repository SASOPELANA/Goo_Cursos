package libro

import "fmt"

type Libro struct {
	Titulo  string
	Autor   string
	Paginas int
}

// Simular Constructor en GO con puntero (Referencia)
func NuevoLibro(titulo string, autor string, paginas int) *Libro {
	return &Libro{
		Titulo: titulo,
		Autor:   autor,
        Paginas: paginas,
	}

}


func (b *Libro) PrintInfo() {
	fmt.Printf("Titulo: %s\nAutor: %s\nPaginas: %d\n", b.Titulo, b.Autor, b.Paginas)
}