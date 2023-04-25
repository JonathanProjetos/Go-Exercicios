package main

import "fmt"

var slice = []int	{}

func main() {
	x:= filterNumberPair(sum,10,10,10,10,10,10,15,15)
	fmt.Println(x)

	y:= filterNumberOdds(sum,11,11,11,11,11,11,24,24)
	fmt.Println(y)
}	

func sum(a ...int)int {
	t:= 0
	for _,v := range a{
		t += v
	}
	return t
}

func filterNumberPair(f func(x ...int)int, p ...int) int{
	for _,v:= range p {
		if v%2 == 0 {
			continue
		}else {
			slice = append(slice, v)
		}
	}
	numbers:= f(slice...)
	return numbers
}

func filterNumberOdds(f func(x ...int)int, p ...int) int{
	for _,v:= range p {
		if v%2 != 0 {
			continue
		} else {
			slice = append(slice, v)
		}
	}
	numbers:= f(slice...)
	return numbers
}


/* 
	- Callback: passe uma função como argumento a outra função.
*/