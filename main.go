package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	return a / b
}

func main() {
	fmt.Println(add(1, 2) + minus(3, 4) + multiply(5, 6) + divide(8, 2))
}
