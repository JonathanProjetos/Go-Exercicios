package main

import (
	"fmt"
)

func main() {
	for num := 33; num <= 122; num++ {
		fmt.Printf("%d, %#x, %#U, \t\n", num, num, num)
		fmt.Printf("%v,\n", string(num))
		//  Pego somente a string referênte ao valor unicode.
	}
}

/* 
- Desafio surpresa!
- Format printing:
    - Decimal       %d
    - Hexadecimal   %#x
    - Unicode       %#U
    - Tab           \t
    - Linha nova    \n
- Faça um loop dos números 33 a 122, e utilize format printing para demonstrá-los como texto/string.
*/