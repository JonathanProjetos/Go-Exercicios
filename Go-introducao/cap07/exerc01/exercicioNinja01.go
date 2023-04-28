package main

import "fmt"

func main() {
	x:= 10
	incrementeValue(x)
}

func incrementeValue(x int) {
	x++
	fmt.Println(x)
	fmt.Println("Endereço de memória:",&x)
}

/* 
	- Crie um valor e atribua-o a uma variável.
	- Demonstre o endereço deste valor na memória.
*/