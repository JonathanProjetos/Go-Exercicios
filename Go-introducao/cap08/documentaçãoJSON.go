package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sort"
)

import "golang.org/x/crypto/bcrypt"

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

	// person05(jamesBond)
	// person06(&info2, json2)

	// sortExample()
	// sortIntsExample()

	//ordenação

	// carros:= []car {
	// 		{"grand siena", 100, 10},
	// 		{"eco esport", 120, 30 },
	// 		{"fusca", 200, 20},
	// }
	// // ordena por nome
	// sort.Sort(ordenarNome(carros))
	// fmt.Println("Nome:\n",carros)

	// // ordena por potência
	// sort.Sort(ordenarPotencia(carros))
	// fmt.Println("Potência:\n",carros)

	// // ordena por consumo
	// sort.Sort(ordenarConsumo(carros))
	// fmt.Println("Consumo:\n",carros)

	// // ordena por lucro do posto
	// sort.Sort(ordenarLucroDoPosto(carros))
	// fmt.Println("Lucro:\n",carros)

	//bcrypt

	senha:= "senhaForte"
	senha2:= "senhaFraca"

	sb, err := bcrypt.GenerateFromPassword([]byte(senha), 10)
	if err != nil {fmt.Println(err)}
	fmt.Println(string(sb))
	//correta return nil	
	fmt.Println(bcrypt.CompareHashAndPassword(sb,[]byte(senha)))
	// errada retorna um erro
	fmt.Println(bcrypt.CompareHashAndPassword(sb,[]byte(senha2)))
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

/* 
	Em Go, uma interface é um tipo abstrato que define um conjunto de métodos. 
	A interface io.Writer é uma das interfaces mais comuns e é usada para representar um objeto que pode escrever dados em algum lugar, 
	como na tela ou em um arquivo. Qualquer tipo que implemente os métodos da interface io.Writer pode ser usado em funções que esperam
	um io.Writer como parâmetro, sem que precise ser especificado o tipo específico. 
	Isso permite que o código seja mais flexível e reutilizável, pois diferentes tipos podem ser usados como entrada para as mesmas funções.
	Além disso, em Go é possível criar interfaces vazias, que não especificam nenhum método, 
	mas podem ser usadas para permitir que um tipo seja usado como argumento genérico em funções e métodos. 
	Isso permite uma maior flexibilidade na escrita de código genérico e reutilizável em Go.
*/


// Pacote sort

//O sort é um pacote para ordenação de slices

func sortStringsExample() {
		ss:= []string{"fui", "jantar", "fora", "hoje"}
		fmt.Println(ss)
		sort.Strings(ss)
		fmt.Println(ss)
}


func sortIntsExample() {
	nn:= []int{10,35,87,9,90,20,58}
	fmt.Println(nn)
	sort.Ints(nn)
	fmt.Println(nn)
}

// Custom Sort

type car struct {
	nome string
	potencia int
	consumo int
}

type ordenarNome []car
type ordenarPotencia []car
type ordenarConsumo []car
type ordenarLucroDoPosto []car

//Por nome

func (n ordenarNome) Len()int {
	return len(n)
}

func (n ordenarNome) Swap(i , j int) {
	n[i], n[j] = n[j], n[i]
}

func (n ordenarNome) Less(i, j int)bool {
	return n[i].nome < n[j].nome
}

//por potência

func (p ordenarPotencia) Len()int {
	return len(p)
}

func (p ordenarPotencia) Swap(i , j int) {
	p[i], p[j] = p[j], p[i]
}

func (p ordenarPotencia) Less(i, j int)bool {
	return p[i].potencia < p[j].potencia
}

//por consumo

func (c ordenarConsumo) Len()int {
	return len(c)
}

func (c ordenarConsumo) Swap(i , j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ordenarConsumo) Less(i, j int)bool {
	return c[i].consumo > c[j].consumo
}

func (l ordenarLucroDoPosto) Len()int {
	return len(l)
}

func (l ordenarLucroDoPosto) Swap(i , j int) {
	l[i], l[j] = l[j], l[i]
}

func (l ordenarLucroDoPosto) Less(i, j int)bool {
	return l[i].consumo > l[j].consumo
}
