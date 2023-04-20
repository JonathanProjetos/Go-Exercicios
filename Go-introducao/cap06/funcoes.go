package main

import "fmt"

var x int
var y int

func main() {
	// basic()

	// k:= oneArgument(10)
	// fmt.Println(k)

	// v:= sum(20,30)
	// fmt.Println(v)

	// c,d:= multipleReturn(10,20,30,40)
	// fmt.Println("O valor da soma é:", c)
	// fmt.Println("O comprimento de x é:", d)

	// a,b,p := exampleParameter("tarde", 10,20,30,40,50)
	// fmt.Println("O valor da soma é:", a)
	// fmt.Println("O comprimento de x é:", b)
	// fmt.Println("O comprimento de x é:", p)
	
	// t:= enumerate(10,20,30,89)
	// fmt.Println(t)

	// examplePrintDefer()
	// nome := pessoa{"jonathan",25}
	// nome.examplePrint()

	// ferramenteiro:= ferramenteiro {
	// 	pessoa1:pessoa1{
	// 		nome: "Jaco",
	// 		sobrenome: "silvim",
	// 		idade: 54,
	// 	},
	// 	especializacao: "torneiaria",
	// 	salario: 5000,
	// }
	
	// desenvolvedor := desenvolvedor {
	// 	pessoa1: pessoa1{
	// 		nome: "Ana",
	// 		sobrenome: "batista",
	// 		idade: 40,
	// 	},
	// 	linguagemFavorita: "javascript",
	// 	anos: 10,
	// }

	// ferramenteiro.print()
	// desenvolvedor.print()

	// serHumano(ferramenteiro)
	// serHumano(desenvolvedor)

	// anonimous()
	// expressionFunc()

	// z:= returnFunc()
	// j:= z(80)
	// fmt.Println(j)

	// d:= onlyPairs(sumNubers, 10,31)
	// fmt.Println(d)

	// e:= onlyOdds(sumNubers, 11,20)
	// fmt.Println(e)

	// g:= i()
	// fmt.Println(g())
	// fmt.Println(g())
	// fmt.Println(g())

	// b:= i()
	// fmt.Println(b())
	// fmt.Println(b())
	// fmt.Println(b())

	n:= recursive(6)
	fmt.Println(n)

	m:= loop(6)
	fmt.Println(m)
}

// Tudo em Go e pass by value direferte da linguagem python como exemplo sendo pass by referênce.
// Em Go, todos os parâmetros de uma função devem ter tipos especificados, assim como o tipo de retorno esperado pela função.

func basic() {
	fmt.Println("olá! bom dia.")
}

func oneArgument(x int)string {
	return fmt.Sprintf("O valor do argumento x é: %d", x)
}


func sum(x,y int) int{
	return x + y
}

//O valor de x pode ser uma inifidade de dados do tipo especificado, neste exemplo passe 4 parâmetros multipleReturn(10,20,30,40).

func multipleReturn(x ...int) (int, int) {
	// função com parâmetros variadicos
	s:= 0 

	for _,v := range x {
		s += v
	}

	return s, len(x)
}

//Os parâmetros variadicos DEVEM sempre vir no final.
func exampleParameter(s string,x ...int)(int,int,string) {
	z := ""
		if s == "manhã" {
			z = "Bom dia!"
		}else if s == "tarde" {
			z = "Boa tarde"
		} else {
			z = "Boa noite"
		}

	soma := 0 

	for _,v := range x {
		soma += v
	}
	
	return soma, len(x), z
}

//Desenrolando o enumerate

func enumerate(param ...int)int {
	sum := 0
	for _,g := range param {
		sum += g
	}

	return sum
} 


// Defer

//é uma palavra-chave da linguagem de programação Go que é usada para adiar a execução de uma função até o final do escopo atual.

func examplePrintDefer() {
	defer fmt.Println("com defer printou depois")
	fmt.Println("sem defer printou primeiro")
	fmt.Println("\n ----")
	//Caso tenha feito mais de 1 uso de defer a sequência de execução sempre será os ultimos seram os primeiros
	//como no exemplo abaixo primero e printado o 3,4 logo depos o 2,1
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)

}

//Métodos

type pessoa struct {
	nome string
	idade int
}

// O método adiciona funcionalidade ao tipo, limitando a função a usar somente os dados do tipo recebido. 

func(p pessoa) examplePrint() {
	fmt.Println(p.nome, "diz bom dia!")
}


//Interfaces & Polimosfismo

type pessoa1 struct {
	nome string
	sobrenome string
	idade int
}

type ferramenteiro struct {
	pessoa1
	especializacao string
	salario int
}

type desenvolvedor struct {
	pessoa1
	linguagemFavorita string
	anos int
}

func(f ferramenteiro) print() {
		fmt.Println("Meu nome é", f.nome,"tenho especialização em", f.especializacao, "minha média salarial é", f.salario, "e ouve só bom dia!")
}

func(d desenvolvedor) print() {
	fmt.Println("Meu nome é", d.nome,"amo a linguagem de programação",d.linguagemFavorita,"é possuo",d.anos,"de experiência", "e ouve só bom dia!")
}

//interface que tem como método a função print.
type gente interface {
	print()
}

//Função serHumano recebe como parâmetro a interface gente que chama os métodos print.
func serHumano(g gente) {
		g.print()
}


//Funções anõnimas

func anonimous() {
	x:= 10

	//função anõnima estou chamando a função e atribuindo o valor de x a ela.
	func (x int)  {
		fmt.Println(x * 10)
	}(x)

}


//Funcão com expressão

func expressionFunc() {
	x:= 10

	y:=	func (x int) int {
		return x * 10
	}

	y(x)

	fmt.Println("O valor de X foi aumentado para", y(x))
}


// Retornando uma função

func returnFunc() func(int)int{
	return func (t int) int {
		return 	t * 10
	}
}


// Calllbacks

func sumNubers(x ...int) int{
	y:= 0

	for _,v := range x {
		y += v 
	}
	return y
}

func onlyPairs(f func(x ...int)int, y ...int) int {
	var slice []int
	for _,v := range y {
		if v %2 == 0 {
			slice = append(slice, v)
		}
	}
	total:= f(slice...)

	return  total
}

func onlyOdds(f func(x ...int)int, y ...int) int {
	var slice []int
	for _,v := range y {
		if v %2 == 1 {
			slice = append(slice, v)
		}
	}
	total:= f(slice...)

	return  total
}


//Closure

func i() func()int {
	x:= 0
	return func() int {
		x++
		return x
	}
}


//Recursividade

func recursive(x int)int {
	if x == 1 {
		return x
	}

	return x * recursive(x-1)
}

func loop(x int) int {
	y:= 1

	for i := x; i >= 1; i-- {
		if i != 0 {
			y *= i
		}
	}

	return y
}	