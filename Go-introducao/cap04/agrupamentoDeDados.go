package main

import "fmt"

func main() {
	// arrayExample()
	// sliceCompounded()
	// sliceForRange()
	// sliceList()
	// sliceDel()
	//appendSlice()
	// makeInfo()
	// sliceOfSlice()
	// mapsInfo()
	mapDel()
}

// array

var x [5]int
var xx = [] int {1,2,3,4,5,6,7,8,9,10,11,12}

// Em Go, estrutura de dado do tipo array terá somente dados internos referênte a tipagem utilzada. 
// Além disso, não é possível alterar o tamanho de um array em Go após sua declaração.

func arrayExample() {
	x[0] = 10
	x[1] = 15
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Printf("%T", x)
}


// Slice; literal composta

func sliceCompounded() {
	// Error: Array não permite uma quantidade de dados maior do que a que foi declarada anteriormente.
	// array := [5] int {1,2,3,4,5,6}
	// fmt.Println(array)
	//Slice
	slice := [] int {1,2,3,4,5,6,7,9}// ...... 
	fmt.Println(slice)
	//Slice permite uma quantidade ilimitada de valores na sua estrutura

	// Busco o item na estrutura de dados slice pelo índeci
	fmt.Println(slice[3])
	// Atribuo um novo valor ao dado recuperado do slice
	slice[3] = 30000
	fmt.Println(slice)

	// No exemplo abaixo vai gerar um erro devido ao slice não ter o comprimento 20. Sendo assim, não é possível atribuir um valor ao índice 20.
	slice[20] = 1

	fmt.Println(slice[20])
	fmt.Println(slice)

	// Se utilizar o append é possivo atribuir valor na estrutura de dado slice lembrando que o 20 não é a posição len e sim um valor adicionado na posição 9.
	slice = append(slice, 20)
}

func sliceForRange() {
	slice := [] string {"banana", "maçâ", "pera", "uva"}
	slice1 := [] int {1,2,3,4,5,6,7,8,9}

	for i, v := range slice {
		fmt.Println("No índice", i, "temos o valor", v)
	}

	total := 0

	for _, t := range slice1 {
		total += t
	}
	fmt.Println(total)
}

// Fatiamneto ou deletenado uma fatia

func sliceList() {
	t:= [] int {1,2,3,4,5,6,7,8,9,10,11,12}

	//Pego do 4 ao 9
	h:= t[3:9]
	//Pego do 1 ao 5
	k:= t[:5]
	//Pego do 8 ao 12
	j:= t[7:]
	//Pego todos os valores
	b:= t[:]

	fmt.Println(h)
	fmt.Println(k)
	fmt.Println(j)
	fmt.Println(b)
}

func sliceDel() {
	//slice compĺeto
	fmt.Println(xx)

	// pego 1, 2
	gg:= xx[:2]
	fmt.Println(gg)

	//pego do 9 ao 12
	dd:= xx[8:]
	fmt.Println(dd)

	//crio uma nova slice partindo da criada
	xx= append(gg,dd ...)
	fmt.Println(xx)
}

func appendSlice() {
	x:= [] int {1,2,3}
	fmt.Println(x)
	y:= [] int {4,5,6}
	fmt.Println(y)

	x = append(x, 7,8,9)
	fmt.Println(x)
	y = append(y, 10,11,12)
	fmt.Println(y)
	//O método append me permite anexiar infinito slices, criando uma nova estrutura slice.
	// Os ... logo no final é semelhante ao espreed operation em javascript, ele espalha os valores de todos os slices dentro de uma nova slice.
	d:= append(x,y ...)
	fmt.Println(d)
}		

//Slice:make

//make([]T, len, cap)
//Make recebe 3 parâmetros:
//1- []T - recebe o slice definido pelo tipo
//2- len - recebe o comprimento do slice
//3- cap - a definição da capacidade do array que deu origem ao slice []T
// Quando o valor do slice []T excede a capacidade, será criado um novo array com um valor
// aproximado do dobro do inicial, podendo gerar um custo computacional adicional e não gera informações visuais.

func makeInfo() {
	// x:= [] int {1,2,3}
	c := make([]int, 5, 10)

	//Esta expressão vai gerar um 0 no final do slice devido ao len(5)
	c[0], c[1], c[2], c[3] = 10,20,30,40

	//Posso adicionar um valor para atribuir ao 0 em slice c
	c[4] = 50

	// Se uso o append aumento o comprimento. 
	c = append(c, 100)

	fmt.Println(c, len(c), cap(c))
}


//slice: multi-dimencional

func sliceOfSlice() {

	//Agrupamento de slice dentro do identificador ss que é a slice principal.

	ss := [][]int {
		{1,2,3},  //índice 0
		{4,5,6},	//índice 1
		{7,8,9},	//índice 2
	}

	fmt.Println(ss)
	//E possivel resgatar uma slice especifica acessando pelo indice len
	fmt.Println(ss[2])
	//Ou pegar um dado expecifico de dentro de alguma slice
	fmt.Println(ss[1][2])
}

// Maps


func mapsInfo() {
	
	x:= map[int]string{
		123:"bolo",
		856:"leite",
		985:"chcolate",
	}
	
	fmt.Println(x,"\n----")
	
	//intero sobre a variavel x é faço a samo todas as k.
	//Observação importante maps não tem ordem e um range usará uma ordem aleatória.
	
	total := 0
	for k, v := range x {
		fmt.Println(k, v)
		total += k	
	} 

	fmt.Println(total)
}

func mapDel() {
	x:= map[int]string{
		123:"bolo",
		856:"leite",
		985:"chcolate",
	}

	fmt.Println(x,"\n----")

	delete(x, 123)

	fmt.Println(x)
}