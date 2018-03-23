package main

func main() {
	println("hello world")
	var ps *Student1 = new(Student1)
	println((*ps).Name)
	println(ps.Name)
	ps1 := new(Student1)
	ps1.Name = "new"
	println(ps1)
	println(ps1.Name)

	var ps2 = Student1{"name"}
	println(ps2.Name)

	ints := make(map[string]int)
	println(ints)

	s := make([]int, 10) // slice with len(s) == 10, cap(s) == 100
	append := append(s, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123, 123)
	println(append)

}

type Student1 struct {
	Name string
}
