package main

import "fmt"

var x int

func main() {
	x = 1990
	
	for {
		if(x <= 2023) {
			fmt.Println(x)
			x++
		} else {
			break
		}
	}
}

/* 
- Crie um loop utilizando a sintaxe: for {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/