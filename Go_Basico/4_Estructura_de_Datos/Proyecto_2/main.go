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
		// Definimos los colores para el fondo de la consola
		bgGreen := "\x1b[42m"
		bgCyan := "\x1b[46m"
		bgRed := "\x1b[41m"
		bgBlue := "\x1b[44m"
		bgYellow := "\x1b[43m"
		bgWhite := "\x1b[47m"
		reset := "\x1b[0m"
		bgMagenta := "\x1b[45m"

		var opcion int
		fmt.Println(bgBlue + "Seleccione una opción:\n",
			"1. Agregar tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
	    	"4. Eliminar tarea\n",	
			"5. Salir" + reset)

		fmt.Print( bgGreen+ "Ingrese una opción: " + reset)
		fmt.Scanln(&opcion)

		switch opcion {
        case 1:
			var t Tarea
            fmt.Print(bgCyan+"Ingrese el nombre de la tarea: "+reset)
            t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre)
            fmt.Print(bgCyan+"Ingrese la descripción de la tarea: "+reset)
			t.descripcion, _ = leer.ReadString('\n')
			t.descripcion = strings.TrimSpace(t.descripcion)
			lista.agregarTarea(t)
			fmt.Println(bgWhite+"Tarea agregada correctamente."+reset)
		case 2:
			var index int
			fmt.Print(bgMagenta+"Ingrese el índice de la tarea que desea marcar como completada: "+reset)
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println(bgMagenta+"Tarea marcada como completada correctamente."+reset)
		case 3:	
			var index int
			var t Tarea
			fmt.Print(bgWhite+"Ingrese el indice de la tarea que desea actualizar:"+reset)
			fmt.Scanln(&index)
			fmt.Print(bgWhite+"Ingrese el nombre de la tarea: "+reset)
            t.nombre, _ = leer.ReadString('\n')
			t.nombre = strings.TrimSpace(t.nombre)
            fmt.Print(bgWhite+"Ingrese la descripción de la tarea: "+reset)
			t.descripcion, _ = leer.ReadString('\n')
			t.descripcion = strings.TrimSpace(t.descripcion)
			lista.editarTarea(index, t)
			fmt.Println(bgYellow+"Tarea actualizado correctamente."+reset)
		case 4:
			var index int
			fmt.Print(bgRed+"Ingrese el índice de la tarea que desea eliminar: "+reset)
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println(bgYellow+"Tarea eliminada correctamente."+reset)
		case 5:
			fmt.Println(bgGreen+"Saliendo del programa."+reset)
			fmt.Println(" ")
			fmt.Println(bgCyan+"§ ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ §"+reset)
			fmt.Println(bgRed+"  ↓ °°° PROGRAMADOR: SERGIO ALEJANDRO SOPELANA °°° ↑  "+reset)
			fmt.Println(bgCyan+"§ ■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ §\n"+reset)
            return	
		default:
			fmt.Println( bgBlue + "Opcion Inválida." + reset)
		}

		
		// Listar todas las tareas
		fmt.Println(bgRed+"Lista de Tareas:"+reset)
		fmt.Println(bgRed + "====================================================================================" + reset)
		for i, t := range lista.tareas {
		fmt.Printf( bgYellow + "%d. %s - %s - Completado: %t\n" + reset, i, t.nombre, t.descripcion, t.completado)
		} 
		fmt.Println( bgRed + "====================================================================================" + reset)	

	}

}