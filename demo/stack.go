package main

func main() {
	stack := new(MyStack)
	stack.push("123")
	stack.push("123")
	stack.push("123")
	stack.push("123")
	stack.push("123")
	stack.push("123")
	stack.push("123")
	println(stack.toString())
	println(stack.contains("3122"))
	println(stack.contains("123"))
	println(len(stack.data))
	println(stack.pop())
	println(len(stack.data))
	stack.clear()
	println(stack.toString())

}

type MyStack struct {
	data []string
}

func (s *MyStack) push(x string) {
	s.data = append(s.data, x)
}

func (s *MyStack) pop() string {
	result := make([]string, len(s.data)-1)
	copy(result, s.data)
	s.data = result
	return s.data[len(s.data)-1]
}

func (s *MyStack) first() string {
	return s.data[0]
}

func (s *MyStack) clear() {
	for i := 0; i < len(s.data); i++ {
		s.data[i] = ""
	}
}

func (s *MyStack) addAll(stack MyStack) {
	list := make([]string, len(s.data)+len(stack.data))
	copy(s.data, list)
	for i := len(s.data); i < len(list); i++ {
		list[i] = stack.data[i-len(s.data)]
	}
	s.data = list
}

func (s *MyStack) contains(key string) (result bool) {
	for _, v := range s.data {
		if v == key {
			result = true
			return
		}
	}
	result = false
	return
}

func (s *MyStack) toString() (result string) {
	if len(s.data) == 0 {
		return
	}
	result = "MyStack:["
	for i, v := range s.data {
		if i != len(s.data)-1 {
			result += v + ","
		} else {
			result += v
		}
	}
	result = result + "]"
	return
}
