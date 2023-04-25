package main

import "fmt"

func main() {
	x:= funcReturnFunc()
	fmt.Println(x())
}

func funcReturnFunc() func()int {
	return func() int {
		return 50 * 100
	}
}

/* 
	- Crie uma função que retorna uma função.
	- Atribua a função retornada a uma variável.
	- Chame a função retornada.
*/