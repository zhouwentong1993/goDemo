package main

import "strings"

func main() {
	s := "/tmp/canal/conf/example"
	index := strings.LastIndex(s, "/")
	println(s[:index])
	command := "cd des"
	split := strings.Split(command, " ")
	targetPosition := split[1]
	println(targetPosition)
}
