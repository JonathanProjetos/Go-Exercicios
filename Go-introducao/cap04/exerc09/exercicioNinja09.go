package main

import "fmt"

func main() {

	x:= map[string][]string{
		"Silva_Marta": []string{
			"cozinhar",
			"correr",
		},
		"jhon_doe": []string {
			"futebol",
			"volêi",	
		},
	}

	x ["santos_Jessy"] = []string { "natação", "dança"}

	for i,v := range x {
		fmt.Println(i)
		for _, j := range v {
			fmt.Println("\t",j)
		}
	}

}

/* 
	- Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
*/