package main

import "fmt"


type pessoa struct {
	name string
	sobreNome string
	idade int
}

type gamer struct {
	pessoa
	jogoFavorito string
}


func main() {

	pessoa1:= pessoa{
		name: "jonathan",
		sobreNome: "silva",
		idade: 33,	
	}
	pessoaStruct(pessoa1)

	gamer1:= gamer{
		pessoa: pessoa{
			name: "Joana",
			sobreNome: "batista",
			idade: 40,
		},
		jogoFavorito: "league of legends",
	}

	gamerStruct(gamer1)

}


func pessoaStruct(p pessoa){
	fmt.Println("Me chamo",p.name,p.sobreNome, "tenho", p.idade, " anos é adoro mangá.")
}

func gamerStruct(g gamer) {
	fmt.Println("Me chamo",g.name,g.sobreNome, "tenho", g.idade, "anos é adoro o jogo", g.jogoFavorito,".")
}

/* 
	- Crie um tipo struct "pessoa" que contenha os campos:
    - nome
    - sobrenome
    - idade
	- Crie um método para "pessoa" que demonstre o nome completo e a idade;
	- Crie um valor de tipo "pessoa";
	- Utilize o método criado para demonstrar esse valor.
*/