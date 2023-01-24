//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
)


func solution() {
	var number int64
	number = 2
	// print number before changing
	fmt.Printf("Исходное число: %d, число в двоичном представлении: %b\n", number, number)
	i := 1
	// create mask for XOR
	mask := 1 << (i)
	// Чтобы инвертировать i-й бит числа можно воспользоваться выражением A = XOR(1 << N)
	fmt.Printf("Новое число: %d, новое число в двоичном представлении: %b\n", number^int64(mask), number^int64(mask))
}

func main() {
	solution()
}