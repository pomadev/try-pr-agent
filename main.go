package main

import "fmt"

func f(n int) int {
	if n == 0 {
		return 0
	}
	return n + f(n-1)
}

func main() {
	fmt.Println(f(10))
}
