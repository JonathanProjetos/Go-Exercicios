package main

import "fmt"

func main() {
	x:= struct {
		n string
		r int
		y map[int]string
		z []string
	}{
		n: "jão",
		r: 358,
		y: map[int]string {
			123: "carne",
			456: "coca-cola",
		},
		z: []string { "feijão", "arroz", "vinagrete", "maionese"},
	}

	fmt.Println(x)
}

/* 
	- Crie e use um struct anônimo.
	- Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/