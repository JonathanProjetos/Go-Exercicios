package main

import "fmt"

func main() {
	x:= closure()
	y:= closure()

	fmt.Println(x(), y())
	fmt.Println(x(), y())
	fmt.Println(x())
	fmt.Println(x(), y())
	fmt.Println(x(), y())
	fmt.Println(x(), y())
}

func closure() func()int {
	x:= 0
	return func() int {
		x++
		return x
	}

}

/* 
	- Demonstre o funcionamento de um closure.
	- Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/