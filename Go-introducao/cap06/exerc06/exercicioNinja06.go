package main

import "fmt"

func main() {
	anonimous()
}

func anonimous() {
	x:=20

	func (x int)  {
		fmt.Println(x * 2)
	}(x)
}

/* 
	- Crie e utilize uma função anônima.
*/