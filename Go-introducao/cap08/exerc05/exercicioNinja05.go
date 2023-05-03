package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type orderLastName []user
type orderAge []user
// type orderSayings []user

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	fmt.Println("Original:\n",users, "\n - ")

	sort.Sort(orderLastName(users))

	fmt.Println("SobreNome:\n",users, "\n - ")

	sort.Sort(orderAge(users))

	fmt.Println("Idade:\n",users, "\n - ")

	orderSayings(users)

	fmt.Println("Citações:")

	for i,v:= range users {
		fmt.Println("\t",i," - ","Meu nome é:",v.First, v.Last, "tenho", v.Age,"anos é gosto de comentar sobre:", v.Sayings)
	}
	// fmt.Println("Provérbios:\n",users)
}

func(l orderLastName) Len()int {
	return len(l)
}

func(l orderLastName) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func(l orderLastName) Less(i, j int) bool {
	return l[i].Last < l[j].Last

}

func(a orderAge) Len()int {
	return len(a)
}

func(a orderAge) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func(a orderAge) Less(i, j int) bool{
	return a[i].Age < a[j].Age

}

func orderSayings(u []user) {
		for _,v := range u {
			sort.Strings(v.Sayings)
		}
}



/* 
	- Partindo do código abaixo, ordene os []user por idade e sobrenome.
	- https://play.golang.org/p/BVRZTdlUZ_
	- Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/