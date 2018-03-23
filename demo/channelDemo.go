package main

func main() {
	a := []int{1,2,41,5,3,-1}
	chanints := make(chan int)
	go sum(a[:len(a)/2], chanints)
	go sum(a[len(a)/2:], chanints)
	x := <- chanints
	y := <- chanints
	println(x)
	println(y)
}



func sum(arr []int, c chan int )  {
	total := 0
	for _, value := range arr {
		total += value
	}
	c <- total
}

