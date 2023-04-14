package main

import "fmt"

var x bool

func main() {
	x = false

	switch {
	case x:
		fmt.Println("Sou verdadeiro")
	default:
		fmt.Println("Sou false")
	}
}

/* 
	- Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/