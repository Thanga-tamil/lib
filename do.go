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

func Do (x, y int, arg string) int {
	
	k := key{a: x, b: y}

	return exec(k, arg)

}

func (k key) sum() int {
	return k.a + k.b
}

func (k key) sub() int {
	return k.a - k.b
}

func exec(a arithmetic, Type string) int {

	fmt.Println("input type: ", Type)

	fmt.Println(a)

	if Type == "add" {
		result := a.sum()
		fmt.Println(result)
		return result
	}
	result := a.sub()
	fmt.Println(result)
	return result
}


