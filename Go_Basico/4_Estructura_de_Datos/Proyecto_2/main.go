package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"github.com/fatih/color"
) 

// Estructura Tares
type Tarea struct {
	nombre string
	descripcion string
	completado bool
}

// Estructura para lista de Tarea
type ListaTarea struct {
	tareas []Tarea
}

// Método para agregar tarea
func (l *ListaTarea) agregarTarea(t Tarea){
	l.tareas = append(l.tareas, t)
}

// Método para marcar como completo
func (l *ListaTarea) marcarCompletado(index int){
	l.tareas[index].completado = true
}

// Método editar tarea
func (l *ListaTarea) editarTarea(index int, t Tarea){
	l.tareas[index] = t
}

// Método para eliminar tarea
func (l *ListaTarea) eliminarTarea(index int){
	l.tareas = append(l.tareas[:index], l.tareas[index + 1:]...)
}



func main() {
	// Intancia de la lista de tareas
	lista := ListaTarea{}


	// Instancia de bufio para entrada de datos
	leer := bufio.NewReader(os.Stdin)
	for {

		// Crear funciones de color
		red := color.New(color.FgRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		blue := color.New(color.FgBlue).SprintFunc()
		yellow := color.New(color.FgYellow).SprintFunc()
		cyan := color.New(color.FgCyan).SprintFunc()
		magenta := color.New(color.FgMagenta).SprintFunc()
		white := color.New(color.FgWhite).SprintFunc()

		var opcion int
		fmt.Println(blue("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
	    	"4. Eliminar tarea\n",	
			"5. Salir"))

		fmt.Print(red("Ingrese una opción: "))
		fmt.Scanln(&opcion)

		switch opcion {
        case 1:
			var t Tarea
            fmt.Print(cyan("Ingrese el nombre de la tarea: "))
            t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre)
            fmt.Print(cyan("Ingrese la descripción de la tarea: "))
			t.descripcion, _ = leer.ReadString('\n')
			t.descripcion = strings.TrimSpace(t.descripcion)
			lista.agregarTarea(t)
			fmt.Println(white("Tarea agregada correctamente."))
		case 2:
			var index int
			fmt.Print(magenta("Ingrese el índice de la tarea que desea marcar como completada: "))
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println(magenta("Tarea marcada como completada correctamente."))
		case 3:	
			var index int
			var t Tarea
			fmt.Print(white("Ingrese el indice de la tarea que desea actualizar:"))
			fmt.Scanln(&index)
			fmt.Print(white("Ingrese el nombre de la tarea: "))
            t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre)
            fmt.Print(white("Ingrese la descripción de la tarea: "))
			t.descripcion, _ = leer.ReadString('\n')
			t.descripcion = strings.TrimSpace(t.descripcion)
			lista.editarTarea(index, t)
			fmt.Println(yellow("Tarea actualizado correctamente."))
		case 4:
			var index int
			fmt.Print(red("Ingrese el índice de la tarea que desea eliminar: "))
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println(yellow("Tarea eliminada correctamente."))
		case 5:
			fmt.Println(green("Saliendo del programa."))
			fmt.Println(" ")
			fmt.Println(cyan("§ ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ §"))
			fmt.Println(red("  ↓ °°° PROGRAMADOR: SERGIO ALEJANDRO SOPELANA °°° ↑  "))
			fmt.Println(cyan("§ ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ §\n"))
			fmt.Println(" ")
			fmt.Println(cyan("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■"))
			fmt.Println(cyan("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■"))
			fmt.Println(white("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■"))
			fmt.Println(white("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■"))
			fmt.Println(cyan("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■"))
			fmt.Println(cyan("■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■"))
			fmt.Println(" ")



            return	
		default:
			fmt.Println(blue("Opcion Inválida."))
		}

		
		// Listar todas las tareas
		fmt.Println(red("Lista de Tareas:"))
		fmt.Println(red("===================================================================================="))
		for i, t := range lista.tareas {
		fmt.Printf( yellow("%d. %s - %s - Completado: %t\n"), i, t.nombre, t.descripcion, t.completado)
		} 
		fmt.Println( red("===================================================================================="))	

	}

}