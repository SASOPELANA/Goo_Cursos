package main

import (
	"fmt"
)

func main() {
	vectorMatriz()
	Rebana()
	Mapas() 
}

func vectorMatriz(){
	var a = [...]int{30, 40, 50, 60, 80, 80}
	a[0] = 20

	for x := 0; x < len(a); x++ {
		fmt.Println(a[x])
	}

	// Iterar el valor y el indice del arreglo o matriz
	for index, value := range a {
		fmt.Printf("Indice: %d, Valor: %d\n", index, value)
	}

	var vector[5] int
	vector[0] = 200 
	vector[1] = 300 
	vector[2] = 400 
	vector[3] = 500
	vector[4] = 600

	for i := 0; i < 5; i++ {
		fmt.Print(vector[i]," ")
	}

	// Matriz bidimensional 
	var matriz = [3][3]int {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println("\n",matriz)
}

func Rebana(){
	// Esto es una rebanada en Go, es similar al arreglo
	diasSeamana := [] string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}

	diaRebana := diasSeamana[0:5] // Rebanada en Go

	fmt.Println(diasSeamana)
	fmt.Println(diaRebana)

	fmt.Println(len(diaRebana)) // Muestra lo que tiene la rebanada -->    len()
	fmt.Println(cap(diaRebana)) // Muestra la capacidad de la rebanada -->    cap()

	// Agregar un nuevo elemento a la rebanada
	diaRebana = append(diaRebana, "Viernes") // Cona la funcion append() se agrega un elemento a la rebanada
	fmt.Println(diaRebana)


	// Se puede aumentar o quitar la capacidad de la rebanada con al funcion append()
	diaRebana = append(diaRebana, "Viernes", "Sabado", "Otros dias mas....", "Otra semana") // Cona la funcion append() se agrega un elementos a la rebanada, se aumenta o se quita.
	fmt.Println(diaRebana)
	fmt.Println(len(diaRebana)) // Muestra lo que tiene la rebanada -->    len()
	fmt.Println(cap(diaRebana)) // Muestra la capacidad de la rebanada -->    cap()

	// Quitar elementos en al rebanada
	diaRebana = append(diaRebana[:2], diaRebana[3:]...)
	fmt.Println("\n",diaRebana)
	fmt.Println(len(diaRebana)) // Muestra lo que tiene la rebanada -->    len()
	fmt.Println(cap(diaRebana)) // Muestra la capacidad de la rebanada -->    cap()

	// Funcion make() para craer rebanadas
	nombres := make([]string, 5, 10)
	nombres[0] = "Juan"
	fmt.Println("\n",nombres) 

	// Funcion para copiar rebanadas. Funcion -->     copy()
	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	copy(rebanada2, rebanada1)
	fmt.Println("\n",rebanada1)
	fmt.Println("",rebanada2)
	fmt.Println(copy(rebanada2, rebanada1)) // Muestra cuantos elementos se copiaron.
}



func Mapas(){
	colores := map[string]string{
		"rojo": "#FF0000",
		"verde": "#00FF00",
		"azul": "#0000FF",
	}

	fmt.Println(colores["rojo"])
	colores["negro"] = "#000000"

	fmt.Println(colores)
	valor, ok := colores["verde"]

	if ok {
		fmt.Println("Si existe la clave")
	}else{
		fmt.Println("No existe la clave")
	}
	fmt.Println(valor)

	valor, ok = colores["blnaco"]
	if ok {
		fmt.Println("Si existe la clave")
	}else{
		fmt.Println("No existe la clave")
	}
	fmt.Println(valor)

	// Eliminar elemnto del mapa. Funcion -->     delete()
	delete(colores, "rojo")
	fmt.Println(colores)

	fmt.Println(" ")

	// Iterar elementos de un mapa
	for clave, valor := range colores {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}
}


