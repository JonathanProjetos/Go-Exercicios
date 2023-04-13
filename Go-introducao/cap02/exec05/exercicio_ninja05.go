package main

import "fmt"

var x string

func main () {
	x = `
	testando\
	raw\
	string\
	literal`
	
	fmt.Println(x)
}

/* 
- Crie uma variável de tipo string utilizando uma raw string literal.
- Demonstre-a.
*/