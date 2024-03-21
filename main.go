package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println(add(1, 2) + add(3, 4) + add(5, 6))
}
