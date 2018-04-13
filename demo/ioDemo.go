package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"os"
)

var commandList = []string{"ls", "cd", "rm", "cp", "mv"}

var (
	command string
	targetPosition  string
	currentLocate = "/Users/finup123"
	beforeLocate  = currentLocate
)

func main() {
	for {
		fmt.Println("enter command!: ")
		fmt.Scanln(&command,&targetPosition)
		fmt.Println("The command you typed is :", command)
		fmt.Println("This command ", If(checkCommand(command), "is valid", "is not valid"))
		switch command {
		case "ls":
			lsLogic(currentLocate)
		case "cd":
			cdLogic()
		case "rm":
			println("rm")
		case "cp":
			println("cp")
		case "mv":
			println("mv")
		}
	}
}
func cdLogic() {
	if targetPosition == ".." {
		if currentLocate == "/" {
			beforeLocate = currentLocate
			currentLocate = "/"
		} else {
			beforeLocate = currentLocate
			currentLocate = currentLocate[:strings.LastIndex(currentLocate, "/")]
		}
	} else if targetPosition == "-" {
		currentLocate, beforeLocate = beforeLocate, currentLocate
	} else {
		tempLocate := currentLocate
		tempLocate = tempLocate + "/" + targetPosition
		stat, err := os.Stat(tempLocate)
		if err != nil {
			return
		}
		if stat.IsDir() {
			currentLocate, beforeLocate = beforeLocate, currentLocate
			tempLocate, currentLocate = currentLocate, tempLocate
		}

	}

	println("Now current location is ", currentLocate)
}

func lsLogic(fileName string) string {
	dir, _ := ioutil.ReadDir(fileName)
	for _, file := range dir {
		if file.IsDir() {
			println(file.Name(), " is a dir")
		} else {
			println(file.Name(), " is a file")
		}
	}
	return ""
}

func checkCommand(command string) (result bool) {
	for _, value := range commandList {
		if strings.HasPrefix(command, value) {
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
