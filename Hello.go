package main

import "fmt"

func HelloWorld() string {
	messageHelloWorld := "Hello World from Go !!"
	return messageHelloWorld
}

func main() {
	HelloWorld()
	fmt.Println("Hello World from Go !!")
}
