package main

import "fmt"

func main() {

	x:= [][]string {
		[]string {
			"jonathan",
			"filho",
			"cozinhar",
		},
		[]string {
			"Marta",
			"Ana",
			"volêi",
		},
		[]string {
			"Jennifer",
			"silva",
			"surf",
		},
	}

	for i, v := range x {
		fmt.Println(i," - ", v[0])
		fmt.Println(i," - ", v[1])
		fmt.Println(i," - ", v[2])
	}
}

/* 
	- Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
  - "Nome", "Sobrenome", "Hobby favorito"
	- Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/