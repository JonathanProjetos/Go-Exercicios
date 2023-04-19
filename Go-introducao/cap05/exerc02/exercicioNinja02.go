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

	x := map[string]pessoa{}

	x["Silveira"] = pessoa1 


	for k,v := range x {
		fmt.Println(k ," - " ,v.nome)
		for _,s := range v.saboresFavoritoDeSorvete{
			fmt.Println("\t",s)
		}
	}

	pessoa2 := pessoa {"Diego", "Silveira", []string{"Morango", "Nata"}}

	y:= map[string]pessoa {}

	y["Silveira"] = pessoa2

	for k,v := range y {
		fmt.Println(k ," - " ,v.nome)
		for _,s := range v.saboresFavoritoDeSorvete{
			fmt.Println("\t",s)
		}
	}

}

/* 
	- Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
	- Demonstre os valores do map utilizando range.
	- Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/