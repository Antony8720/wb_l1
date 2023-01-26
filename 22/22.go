package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewFloat(10000000)
	b := big.NewFloat(20000000)

	addition := new(big.Float).Add(a, b)
	multiplication := new(big.Float).Mul(a, b)

	subtraction := new(big.Float).Sub(a, b)
	division := new(big.Float).Quo(a, b)

	fmt.Printf("Add: %1.f\n", addition)
	fmt.Printf("Multiply: %1.f\n", multiplication)

	fmt.Printf("Discard: %1.f\n", subtraction)
	fmt.Println("Devide:", division)
}