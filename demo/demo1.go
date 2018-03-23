package main

import (
	"fmt"
)

func Bob(s Student) {
	s.Name = "Bob" // changes only the local copy
}

// Charlie sets pp.Name to "Charlie".
func Charlie(ps *Student) {
	ps.Name = "Charlie"
}

type Student struct {
	Name string
}

func main() {

	s := Student{"Alice"}

	Bob(s)
	fmt.Println(s) // prints {Alice}

	Charlie(&s)
	fmt.Println(s) // prints {Charlie}
	c, d := func(a int, b int) (c int, d int) {
		c = a + b
		d = 4
		return
	}(1, 2)
	fmt.Println(c, d)

	e := func(a int, b int) (c int) {
		return a + b
	}(1, 2)
	fmt.Println(e)

	i := 1
	point := &i    // point等于一个地址
	data := *point //data == 1
	println(data)
	println(point)
	type SmallSoho struct {
		Name string //公有变量大写
		Sex  string
		kg   int //私有变量小写
	}
	meStruct := new(SmallSoho)
	println(meStruct)

	meStruct1 := &SmallSoho{"small", "男", 1} //可以初始化
	println(meStruct1)
	println(meStruct1)

	meStruct2 := SmallSoho{}
	meStruct2.kg = 23
	println(meStruct2.kg)

}

type Stack struct {
	data []string
}

// Push adds x to the top of the stack.
func (s *Stack) Push(x string) {
	s.data = append(s.data, x)

}
