package main

import "fmt"

var x int

func main() {
	x = 1990

	for x <= 2023 {
		fmt.Println(x)
		x++
	}
}

/* 
- Crie um loop utilizando a sintaxe: for condition {}
- Utilize-o para demonstrar os anos desde que você nasceu.
*/