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

	for i,v := range x {
		fmt.Println(i)
		for _, j := range v {
			fmt.Println("\t",j)
		}
	}

}


/* 
	- Crie um map com key tipo string e value tipo []string.
    - Key deve conter nomes no formato sobrenome_nome
    - Value deve conter os hobbies favoritos da pessoa
	- Demonstre todos esses valores e seus índices.
*/