package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
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
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
	    	"4. Eliminar tarea\n",	
			"5. Salir")

		fmt.Print("Ingrese una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
        case 1:
			var t Tarea
            fmt.Print("Ingrese el nombre de la tarea: ")
            t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre)
            fmt.Print("Ingrese la descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			t.descripcion = strings.TrimSpace(t.descripcion)
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente.")
		case 2:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea marcar como completada: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente.")
		case 3:	
			var index int
			var t Tarea
			fmt.Print("Ingrese el indice de la tarea que desea actualizar:")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
            t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre)
            fmt.Print("Ingrese la descripción de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			t.descripcion = strings.TrimSpace(t.descripcion)
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizado correctamente.")
		case 4:
			var index int
			fmt.Print("Ingrese el índice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente.")
		case 5:
			fmt.Println("Saliendo del programa.")
            return	
		default:
			fmt.Println("Opcion Inválida.")
		}

		
		// Listar todas las tareas
		fmt.Println("Lista de Tareas:")
		fmt.Println("====================================================================================")
		for i, t := range lista.tareas {
		fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)
		} 
		fmt.Println("====================================================================================")	

	}

}