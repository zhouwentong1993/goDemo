package main

import "fmt"

var commandList = []string{"ls", "cd", "rm", "cp", "mv"}

var (
	command, s string
	currentLocate = "/Users/finup123"
	i          int
	f          float32
	input      = "56.12 / 5212 / Go"
	format     = "%f / %d / %s"
)

func main() {
	for {
		fmt.Println("enter command!: ")
		fmt.Scanln(&command)
		fmt.Println("The command you typed is :", command)
		fmt.Println("This command ", If(checkCommand(command), "is valid", "is not valid"))
		switch command {
		case "ls":
			
			println("ls")
		case "cd":
			println("cd")
		case "rm":
			println("rm")
		case "cp":
			println("cp")
		case "mv":
			println("mv")
		}
		fmt.Sscanf(input, format, &f, &i, &s)
		fmt.Println("From the string we read: ", f, i, s)
	}
}

func lsLogic() string{

	return ""
}

func checkCommand(command string) (result bool) {
	for _, value := range commandList {
		if value == command {
			result = true
			return
		}
	}
	return
}

func If(b bool, t, f interface{}) interface{} {
	if b {
		return t
	}
	return f
}
