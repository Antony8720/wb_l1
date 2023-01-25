package main

import (
	"fmt"
)

func solution() {
	a, b := 1, 2
	fmt.Printf("a: %v, b: %v\n", a, b)
	a, b = b, a 
	fmt.Printf("a: %v, b: %v\n", a, b)
}

func main() {
	solution()
}