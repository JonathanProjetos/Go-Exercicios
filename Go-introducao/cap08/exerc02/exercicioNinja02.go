package main

import (
	"fmt"
	"encoding/json"
)

type person struct {
	First   string   
	Last    string   
	Age     int      
	Sayings []string 
}

var human []person

func main() {
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	convertForGo(&human, []byte(s))

	fmt.Println(human[1].Age)
}

func convertForGo(p *[]person, x []byte ) {
	err := json.Unmarshal(x, &p)
	if err != nil {
		fmt.Println("error:",err)
	}
}


/* 
	- Partindo do código abaixo, utilize unmarshal e demonstre os valores.
	- https://play.golang.org/p/BVRZTdlUZ_
*/