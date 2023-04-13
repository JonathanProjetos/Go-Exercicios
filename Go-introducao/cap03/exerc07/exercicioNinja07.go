package main

import "fmt"

var x int

func main() {
	x = 11
	
	if x == 10 {
		fmt.Println("igual")
	} else if x < 10 {
		fmt.Println("menor")
	} else {
		fmt.Println("maior")
	}
}

/* 
- Utilizando a solução anterior, adicione as opções else if e else.
*/