package main

import "fmt"

func main() {
	// loopFor()
	// loopHours()
	// infityLoop()
	// loopAndContinue()
	// conditional()
	// conditional2()
	// conditionalSwitch()
	// conditionalSwitchType()
	conditionalOperator()
}

var x int 
var y = [10]string{"a","b","c","d","e","f","g","h","i","j"}
// var z string

func loopFor () {
	for i := 0; i < len(y); i++ {
		fmt.Println(i,"-",y[i])
	}

}

func loopHours() {
	for horas := 1; horas <= 12; horas++ {
		fmt.Println("Horas:", horas)
		for min := 1; min <= 60; min++ {
			fmt.Println("Minutos:", min)
			for sec := 1; sec <= 60; sec++ {
				fmt.Println("Segundos:", sec)
			}
		}
	}
}

// Em Go não tem While

func infityLoop() {
	x = 0
	for{
		if x < 10 {
			x++
			fmt.Println("Sou menos que dez")	
		} else {
			fmt.Println("Sou maior que dez")
			break
		}
	}
}


func loopAndContinue() {
	x = 0
	for x <= 20 {
		x++
		if (x%2 != 0) {
			// Se for impar desconsidere o valor é continue.
			continue
		}
		fmt.Println(x)
	}
}

// O continue serve para quebrar a interação no momendo de uma condição.
// O break para a execução do loop.

func conditional() {
	x = 10

	if x == 10 {
		fmt.Println("sou igual")	
	} else {
		fmt.Println("Não sou igual")
	}

	//Negação

	if(x != 10) {
		fmt.Println("Não sou diferente de dez")
	} else {
		fmt.Println("Sou iqual a dez")
	}
}


func conditional2() {
	x = 100
	if(x > 1000) {
		fmt.Println("Sou maior que 1000")
	} else if (x == 100) {
		fmt.Println("Sou igual a 100")
	} else {
		fmt.Println("Sou menor que 1000")
	}
}

func conditionalSwitch() {
	x = 100

	// O switch procura de cima para baixo a primeira condição que satisfaça o caso.

	switch true {
		case x > 1000:
			fmt.Println("Sou maior que 1000")
		case x < 1000:
			fmt.Println("Sou menor que 1000")
		case x == 1000:
			fmt.Println("Sou igual a 1000")
	}

	//condicional switch usando o fallthorough

	pessoa:= "Jonathan"

	switch pessoa {
		case "Marcos":
			fmt.Println("Hoje quem está no parque é Marcos")
			fallthrough
		case "João":
			fmt.Println("Quando Marcos está no parque, joão tambem estará")
		case "Tatiane":
			fmt.Println("Hoje quem está no parque é Tatiane")
			fallthrough
		case "Ana":
			fmt.Println("Quando Tatiane está no parque, Ana tambem estará")
		default:
			fmt.Println("O parque está fechado")
	}


	switch pessoa {
		case "Leo", "Marta":
			fmt.Println("Adora pizza de queijo")
		case "Sabrina", "Jennifer":
			fmt.Println("Adora pizza de calabresa")
		default:
			fmt.Println("Não gosta de pizza")
	}
}

var u interface{}

func conditionalSwitchType() {
		u = -20

	switch u.(type){
		case bool:
			fmt.Println("É um boleano")
		case string:
			fmt.Println("É uma string")
		case int:
			fmt.Println("É um número")
		case float64:
			fmt.Println("É um número flutuante")
		default:
			fmt.Println("Tipo não mapeado")
			
		}
}

func conditionalOperator() {
	l := 103

	if l <= 99 || l > 102  {
		fmt.Println("Sou menor igual a 100 ou maior que 102")
	} else if l == 100 && l != -110 {
		fmt.Println("Sou igual a 100 é diferente de -110")
	} else {
		fmt.Println("Não mapeado")
	}
}