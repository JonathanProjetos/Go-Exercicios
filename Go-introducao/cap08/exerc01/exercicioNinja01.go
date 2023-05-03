package main

import (
	"fmt"
	"encoding/json"
)
type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "James",
		Age:   32,
	}

	u2 := user{
		First: "Moneypenny",
		Age:   27,
	}

	u3 := user{
		First: "M",
		Age:   54,
	}

	users := []user{u1, u2, u3}

	person(users)

}

func person(u []user) {
	v,err := json.Marshal(u)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(v))
}

//Depois de alguns preciosos minutos, lembre-se que para fazer a importação, a estrutura do struct deve estar escrita em letra maiúscula.

/* 
	- Partindo do código abaixo, utilize marshal para transformar []user em JSON.
	 - https://play.golang.org/p/U0jea43X55
*/