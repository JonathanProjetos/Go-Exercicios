package main

import "fmt"

var x= [] int {10,20,30,40,50,60}

func main() {
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5], "\n----")

	//Ou

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}

/* 
- Tente acessar todos os índices de um slice sem usar range.
*/