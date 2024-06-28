package main

import "fmt"

func main() {
	i := 1

	// Tipo while
	for i <= 10 {
		fmt.Println(i)
		i++
	}
	
	// For estandar
	fmt.Print("\n")
	for x := 1; x <= 10; x++{
		if x <=9 {
			fmt.Print(x, ", ")
		}else{
			fmt.Print(x,".")
		}
		
	}

	// continue    salta a la siguiente iteracion
	// break       cancela el ciclo
}