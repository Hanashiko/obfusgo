package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	msg := "Hello, world!"
	fmt.Println(msg)
	fmt.Println("2+3 =", add(2, 3))
}
