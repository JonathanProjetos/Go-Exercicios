package main

import "fmt"

func main() {
	// structExample()
	// nestedStructure()
	anonymousStructs()
}

// Os structs petimite o armazenamento de tipos de dados direntes na mesma estrutura.

type cliente struct {
	nome string
	sobreNome string
	altura float64
	telefone int
	fumante bool
}

func structExample() {
	cliente1 := cliente{
		nome: "jonathan",
		sobreNome: "pereira",
		altura: 1.85,
		telefone: 3199999999,
		fumante: true,
	}

	cliente2 := cliente {"jennifer", "silva", 1.50, 3188898989, false}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}


// Structs embutidos

type pessoa struct {
	name string
	idade int
}

type profissional struct {
	pessoa
	titulo string
	salario int
}

// Nas estruturas structs é possivel acessar conteúdos especificos utilizando Dot notation (.).

func nestedStructure() {
	pessoa1 := pessoa {
		name: "Jessica",
		idade: 40,
	}
	fmt.Println(pessoa1)

	pessoa2 := profissional {
		pessoa: pessoa{
			name: "Jorge",
			idade: 20,
		},
		titulo: "Desenvolvedor",
		salario: 10000,
	}
	
	fmt.Println(pessoa2.name)
	fmt.Println(pessoa2.pessoa.idade) // 20

	//como a estrutura não possui outra chave idade eu posso acessa-la diretamente.
	fmt.Println(pessoa2.idade) // 20

	//Criação de uma struct composta de forma consisa

	pessoa3 := profissional{pessoa{"Juca", 40}, "Desenvolvedor",20000}

	fmt.Println(pessoa3)

}

//Structs anñimos

func anonymousStructs() {
	y:= struct{
		name	string
		idade int
	}{
		name: "Jerson",
		idade: 15,
	}

	fmt.Println(y)	
}