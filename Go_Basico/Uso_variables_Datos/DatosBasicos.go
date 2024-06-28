package main

import (
	"fmt"
	// "strconv" // Uso de esta libreria para conventir un string a un int
	"math" // Para calculos matematicos complejos
) 

func main() {
	fmt.Println("Hola")

	// int  --> Almacena numeros positivos y negativos y con bits dependiendo la pc. Ejemplo 64 bits
	// int8 --> Almacena numeros entre el -128 y 127
	// uint  --> Almacena numeros positivos de cero en adelante

	// float --> Almacena numeros decimales, postivos y negativos
	// float32 --> Almacena numeros
	// ufloat --> Almacena numeros positivos de cero en adelante
	// ----------------------------------------------------------------
	// Boleanos
	/*
	bandera := true
	fmt.Println(bandera)

	fullName := "Sergio Alejandro \t(Alias \"sope\")\n"
	fmt.Println(fullName)

	// Variables tipo byte, para almacenar codigos ascii y caracteteres
	var a byte = 'a'
	fmt.Println("Valor del código Ascii a: ", a) // Muestra el valor ascci el número 97

	// Accder al valor decimal de una cadena. Ejemplo
	s := "Hola"
	fmt.Println("Valor decimal de primer indice de Hola (H): ", s[0])

	// Valor unico. Se hace uso de la palabra reservada rune
	var r rune = '❤'
	fmt.Println("Valor unicode del simbolo ❤ : ", r)
	*/
	// ---------------------------------------------------------------
	// Varibles y su valores predertimado o por defecto. Valor cero
	/*
	var (

		defaulInt int
		defaulUint uint
		defaulFloat float32
		defaulBool bool
		defaulString string
	)

	fmt.Println("Valor por defecto de int, uint, float, bool, string: ", defaulInt, defaulUint, defaulFloat, defaulBool, defaulString)
	*/
     
	// ---------------------------------------------------------------
	// Conversion de tipos de datos
	/*
	var interger16 int16 = 50
	var interger32 int32 = 100
	// fmt.Println(interger16 + interger32) // Esto no se puede hacer. Se debe hacer un conversion de datos primero
	fmt.Println(int32(interger16) + interger32)

	// Conversion de string a entero // Uso de strconv.Atoi  --> de string a entero
	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println(i + i)

	// Conversion de entero a string // Uso de strconv.Itoa  --> de entero a string
	n := 42
	s =  strconv.Itoa(n)
	fmt.Println(s + s)
	*/

	// ---------------------------------------------------------------
	// El paquete fmt
	/*
	fmt.Print("Otro mensaje\n") // Print no hace salto de linea como Println
	// Con Printf podemos formatear informacion 

	var name string
	var age int

	// Pedir datos en Go
	fmt.Print("Ingrese su nombre: ")
	fmt.Scanln(&name)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)	

	greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.\n", name, age)
	fmt.Println(greeting)

	fmt.Printf("El tipo de name es: %T\n", name)
	fmt.Printf("El tipo de age es: %T\n", age)
	*/

	// ----------------------------------------------------------------
	// Operadores aritmeticos y paquete Math
	// Contador en Go x++ 
	// Decremento en Go x--
	// Acumulador en Go  +=     
	// Resta en acumulacion -=

	a := 10
	b := 3

	fmt.Println(a / b)
	fmt.Println(math.Pi)
	fmt.Println(math.Sqrt(64)) // Para sacar la raiz cuadrada
	fmt.Println(math.Cbrt(27)) // Para sacar la raiz cubica
}