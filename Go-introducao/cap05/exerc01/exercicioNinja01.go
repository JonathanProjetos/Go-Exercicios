package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	saboresFavoritoDeSorvete []string
}

func main() {
	pessoa1 := pessoa {
		nome: "Clarice",
		sobrenome: "Silveira",
		saboresFavoritoDeSorvete: []string{"Creme", "Chocolate"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)

	for _,y := range pessoa1.saboresFavoritoDeSorvete {
		fmt.Println("\t",y)
	}


	pessoa2 := pessoa {"Diego", "Silveira", []string{"Morango", "Nata"}}

	fmt.Println(pessoa2.nome, pessoa2.sobrenome)

	for _,t := range pessoa2.saboresFavoritoDeSorvete {
		fmt.Println("\t",t)
	}
}

/* 
	- Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
    - Nome
    - Sobrenome
    - Sabores favoritos de sorvete
	- Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
*/