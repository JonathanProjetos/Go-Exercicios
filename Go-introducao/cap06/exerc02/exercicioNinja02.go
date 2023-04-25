package main

import "fmt"

var slice = []int {1,2,3,4,5,6,7,8,9,10} 


func main() {
	x:= variadicFunc(1,2,3,4,5,6,7,8,9)
	fmt.Println("print - 1  ",x)

	y:= sliceReceved(slice)
	fmt.Println("print - 2  ",y)
}


func variadicFunc(x ...int)int {
	y:= 0
	for _,v := range x {
		y+=v
	}
	return y
}

func sliceReceved(y []int)int {
	x:= 0
	for _,v := range y {
		x+=v
	}
	return x
}


/* 
- Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
- Passe um valor do tipo slice of int como argumento para a função;
- Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
- Passe um valor do tipo slice of int como argumento para a função.
*/