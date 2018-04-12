package main

func main() {
	ints := make(chan int, 3)
	ints <- 3
	ints <- 3
	ints <- 3
	ints <- 3
	println(ints)
}

func consume() {

}
