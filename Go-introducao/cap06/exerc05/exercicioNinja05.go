package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado int
}

type circulo struct {
	raio float64
}

type figura interface {
	area() float64
}

func main() {
	quadrado:= quadrado {
		lado: 50,
	}

	circulo:= circulo {
		raio: 40.0,
	}

	x:= info(quadrado)
	fmt.Println("quadrado",x)

	y:= info(circulo)
	fmt.Println("circulo",y)

}

func (c circulo) area() float64 {
	x:= 2 * math.Pi
	return x*c.raio
}

func (q quadrado) area() float64 {
	x:= float64(q.lado*q.lado)
	return x
}

func info(f figura) float64 {
 	return f.area()
}

/* 
- Crie um tipo "quadrado"
- Crie um tipo "círculo"
- Crie um método "área" para cada tipo que calcule e retorne a área da figura
    - Área do círculo: 2 * π * raio
    - Área do quadrado: lado * lado
- Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
- Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
- Crie um valor de tipo "quadrado"
- Crie um valor de tipo "círculo"
- Use a função "info" para demonstrar a área do "quadrado"
- Use a função "info" para demonstrar a área do "círculo"
*/