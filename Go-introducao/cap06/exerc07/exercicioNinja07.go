package main

import "fmt"

func main() {
	result(10,58)
}

func result(x int, z int) {
	y:= func(a int, b int)int {
		return a + b
	}

	y(x,z)

	fmt.Println("O Resultado é",y(x,z))
}

/* 
	- Atribua uma função a uma variável.
	- Chame essa função.
*/