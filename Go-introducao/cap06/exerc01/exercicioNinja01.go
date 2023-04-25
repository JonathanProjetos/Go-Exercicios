package main

import "fmt"

func main() {
	x:= returnInt()
	fmt.Println(x)

	y,z:= returnIntAndString()
	fmt.Println(y," - ",z)

}

func returnInt()int {
	x:= 1
	y:= 2

	return x+y
}

func returnIntAndString()(int,string) {
	x:= 1
	y:= 2
	z:= "Este é o resultado."

	return (x+y), z
}



/* 
  - Crie uma função que retorne um int
  - Crie outra função que retorne um int e uma string
  - Chame as duas funções
  - Demonstre seus resultados
*/