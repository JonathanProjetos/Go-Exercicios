package main

import "fmt"

func main() {
	x:= 10
	// addressVar()

	//Aqui foi feita uma cópia de 'x' onde, na função, o valor é atribuído. No entanto, fora da função, o valor continua sendo o valor padrão antes do incremento.
	//receivedValue(x)
	//fmt.Println(x)

	//Aqui, por ser atribuído diretamente no endereço de memória, não é gerada uma cópia e o valor é alterado dentro e fora da função.
	receivedPointer(&x)
	fmt.Println(x)
}

// Ponteiro e uma variável que armazena o endereço de memória de outra variável.
//É possivel acesar o endereço da variável utilizando o simbolo &.

func addressVar() {
	x:=10
	//Estou atribuindo o endereço da memória da variável x a variável y
	y:= &x

	//Print normal
	fmt.Println(x)
	// Print do endereço da memória.
	fmt.Println(&x)
	//Print de y que tem o valor atribuido do endereço de memória de x.
	fmt.Println("y",y)

	//DEREFERENCING
	/* 
		Em Go, referências diretas (ou "direct references") se referem a valores que apontam diretamente para um endereço de memória,
		permitindo que os dados sejam acessados e modificados diretamente. Em outras palavras, 
		uma referência direta é uma variável que contém o endereço de memória de outra variável.
	*/
	//Nete caso y armazena o endereço de memória de x utilizando o símbolo * posso acessar o valor neste endereço de memória.

	fmt.Println("endereço", y)
	fmt.Println("valor dentro do endereço", *y)

	//Exmeplo de retorno
	fmt.Printf("%T \n", x) //int
	fmt.Printf("%T \n", y) //*int ponteiro

	//Exemplo de atribuição de valor no endereço de memória
	z:= 0
	v:= &z
	fmt.Println("Normal",z)
	*v++
	fmt.Println("Incrementado",z)
}

//Quando usar ponteiros

//Quando estou fazendo manipulação de grandes quantidades de dados, pensando em performance, 
//é mais eficiente mandar uma função buscar informações no endereço da memória do que fazer 
//cópias dos valores para serem utilizados.

//Exemplos

func receivedValue(x int) {
	x++
	fmt.Println(x)
}

func receivedPointer(x *int) {
	*x++
	fmt.Println(*x)
}
