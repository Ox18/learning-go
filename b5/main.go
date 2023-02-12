package main

import "fmt"

func normalFunction(message string) {
	fmt.Println(message)
}

func tripeArgument(a int, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func main() {
	normalFunction("Hola mundo!")
	tripeArgument(1, 2, "Hola mundo!")
	fmt.Println(returnValue(2))
}
