package main

import "fmt"


var x int

func main() {
	x = 300
	fmt.Printf("%d\n%b\n%#x\n", x, x, x)
	y := x << 1
	fmt.Println(y)
	fmt.Printf("%d\n%b\n%#x\n", y, y, y)
}


/* 
- Crie um programa que:
    - Atribua um valor int a uma variável
    - Demonstre este valor em decimal, binário e hexadecimal
    - Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    - Demonstre esta outra variável em decimal, binário e hexadecimal
*/