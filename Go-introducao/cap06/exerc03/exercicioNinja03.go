package main

import "fmt"

func main() {
	x:= "vou iniciar em primeiro"
	defer fmt.Println(x)
	
	y:= "vou iniciar em segundo"
	fmt.Println(y)
}

/* 
	- Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/