package main

func main() {
	println(1)

	var pm = new(MyType)
	println(pm)
	var n = pm.Get()
	println(n)
}

type MyType struct{ i int }

func (p *MyType) Get() int {
	return p.i
}
