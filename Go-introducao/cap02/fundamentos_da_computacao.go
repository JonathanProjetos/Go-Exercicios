// Tipos Boleanos
package main

import (
	"fmt"
	"runtime"
)

var x bool

func main () {
	fmt.Println(x)
	x = true
	fmt.Println(x)
	x = (10 > 100)
	fmt.Println(x)
	y := (10 < 100)
	fmt.Println(y)
	t := (10 == 100)
	fmt.Println(t)
	w := (50 != 200)
	fmt.Println(w)

	computer()
	tipos_numéricos()
	// tipos_de_string()
	// sistemas_numericos()
	constantes()
	iotas()
	bite()
}

// como os computadores funcionam

func computer () {
	u:= 750000 / 8
	fmt.Println(u) 

	a := []byte("e")
	b := []byte("é")
	c := []byte("ola")


	fmt.Printf("%v\n, %v\n, %v\n", a, b, c)

}

// Tipos númericos


func tipos_numéricos () {
	yu := 10
	yo := 100.00

	fmt.Printf("%v, %T\n", yu,yu)
	fmt.Printf("%v, %T\n", yo,yo)

	//verifico a imagem do sistema opercional
	fmt.Println(runtime.GOOS)
	// verifico a arquitetura
	fmt.Println(runtime.GOARCH)

	// por padrão sempre é aconselhavel usar o float64 para numeros flutuantes.
	// Ao utilizar o tipo int, é aconselhavem deixar o computador escolher uma versão mais apropriada ex: int32, int64
}

// Tipos de estring


func tipos_de_string () {
	// Strings em Go são imutáveis. Se precisar alterar uma string, é preciso criar uma nova variável adicionando a alteração.
	v := "testando"

	i := `
		testando 
		uma 
		formatação 
		diferente`
	
	g := []byte(i)
	
	fmt.Printf("%v, %T\n", v, v)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Println(g)

	for _, v:= range g {
		fmt.Printf("%v - %T - %#U - %#X\n", v, v, v, v)
	}
	
	for p := 0; p < len(i); p++ {
		fmt.Printf("%v - %T - %#U - %#X\n", i[p], i[p], i[p], i[p])
	}
}

func sistemas_numericos () {
	// binario 2 números
	// 0 - 1 
	// decimal
	// 0 - 9
	// hexadécimal 16 números
	// 1 2 3 4 5 6 7 8 9 a b c d e f

	value := 31337

	fmt.Printf("%b - %d - %x\n" ,value, value, value)

}

const m = "Oi, Bom dia!"
const p = 100.0 
// Alternativa de declaração de constantes
const (alternativa = 10)
//ou
const (
	h = 100
	n = 45
	a = 96
)

var k = "Oi, Bom dia!"

func constantes () {
	// constantes são variaveis imutaveis.
	// As constantes em Go não recebem tipagens no momento de uso. As variáveis recebem tipos no momento da declaração.

	fmt.Printf("%v,%T\n", k,k)
	fmt.Printf("%v,%T\n", m,m)
	fmt.Printf("%v,%T\n", p,p)
}


// IOTA

const (
	zz = iota 
	_= iota
	dd = iota
	aa = iota 
	ff = iota
)

const (
	gg = iota * 10
	jj
	_
	qq
	mm
)


func iotas () {
	// iotas são sucessivas constantes não tipadas inteiras.
	// Caso de uso para quando você não está preocupado com o valor das constantes, mas sim que elas sejam diferentes.	
	// E possǘel descarta valores usando o _ 
 	// É possível atribuir um valor à constante utilizando operadores. Cada referência de constante terá seu valor de posição interagindo com o operador.
	fmt.Printf("%v,%v,%v, %v\n",zz,dd, aa,ff)
	fmt.Printf("%v,%v,%v, %v",gg,jj,qq, mm)
}


// Deslocamento de bites


func bite () {
	// bite operations
	fmt.Print(gg,jj,qq,mm)
	// 0110000
	// 0011000
	// 0001100
	// 0000110
	// 0000011

	//Operador para delocamento de bites >> <<. O delocamento será em função da sua direção.

	xx := 1
	yy := xx << 11
	hh := xx >> 10

	fmt.Printf("%b\n", xx)
	fmt.Printf("%b\n", yy)
	fmt.Printf("%b\n", hh)

}