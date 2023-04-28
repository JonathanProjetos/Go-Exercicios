package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

var jamesBond = pessoa {
	Name: "James",
	Sobrenome: "Bonde",
	Idade: 30,
	Profissao: "heroi",
	Conta: 10000000,
}

var shaco = pessoa {
	Name: "Shaco",
	Sobrenome: "box",
	Idade: 100,
	Profissao: "dublê",
	Conta: 20000000,
}

type pessoa struct {
	Name string
	Sobrenome string
	Idade int
	Profissao string
	Conta int
}

var info pessoa
var info2 pessoa
var json1 = []byte(`{"Name":"James","Sobrenome":"Bonde","Idade":30,"Profissao":"Heroi","Conta":10000000}`)
var json2 = `{"Name":"Shaco","Sobrenome":"Box","Idade":100,"Profissao":"Duble","Conta":20000000}`

// Informação importante: para todos os campos serem convertidos para JSON o valor da chave deve iniciar em maiúsculo.

func main() {
	// person1(jamesBond)
	// // person2(shaco)

	// person3(&info,json1)
	// fmt.Printf("%+v \n",info)

	// person4(&info2, json2)
	// fmt.Printf("%+v",info2)

	person05(jamesBond)
	person06(&info2, json2)
}

/* 
	JSON = javascript object notetion.
*/


// JSON marshal(ordenação)
// o pacote json reponde com dois tipos de dados o valor é o error.


func person1(p pessoa) {
	//convertendo o código de Go para JSON
	v,err := json.Marshal(p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(v))
}

func person2(p pessoa) {
	//convertendo o código de Go para JSON
	v,err := json.Marshal(p)
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(string(v))
}


//Convertendo JSON para Go

func person3(p *pessoa, x []byte) {
	//Convertendo de JSON para GO.
	err := json.Unmarshal(x, &p)
	if err != nil {
		fmt.Println("error:",err)
	}
}

func person4(p *pessoa, x []byte) {
		//Convertendo de JSON para GO.
	err := json.Unmarshal(x, &p)
	if err != nil {
		fmt.Println("error:",err)
	}
}

func person05(p pessoa) {
	//Convertendo dados Go em JSON usando o pacote encoding/json.
	encoding:= json.NewEncoder(os.Stdout)
	if encoding != nil {
		fmt.Println(encoding)
	}
	encoding.Encode(p)
}

func person06(p *pessoa, x string) {
	//Convertendo dados Go em JSON usando o pacote encoding/json.
	reader:= strings.NewReader(x)
	decoding:= json.NewDecoder(reader)

	if decoding != nil {
		fmt.Println(decoding)
	}
	decoding.Decode(&p)
}


// Interface whiter

