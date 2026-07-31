package do

import (
	"fmt"
)

type arithmetic interface{
	sum() int
	sub() int
}
type key struct {
	a, b int
}

func Do (x, y int, arg string) {
	
	k := key{a: x, b: y}

	exec(k, arg)

}

func (k key) sum() int {
	return k.a + k.b
}

func (k key) sub() int {
	return k.a - k.b
}

func exec(a arithmetic, Type string){
	fmt.Println("input type: ", Type)

	fmt.Println(a)

	if Type == "add" {
		fmt.Println(a.sum())
	}else {
		fmt.Println(a.sub())
	}
}


