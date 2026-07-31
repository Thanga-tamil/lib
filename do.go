package do

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
	ops := map[string]func() int{
		"+": a.sum,
		"-": a.sub,
	}
	if f, ok := ops[Type]; ok {
		return f()
	}
	return -1
}
