package main

import "fmt"

var esporteFavorito string

func main() {

	esporteFavorito = "vólei"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Adoro futebol")
	case "vólei":
		fmt.Println("Adoro vólei")
	case "basquete":
		fmt.Println("Adoro basquete")
	default:
		fmt.Println("Tipo não mapeado")
	}
}

/* 
	- Crie um programa que utilize a declaração switch, 
	onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/