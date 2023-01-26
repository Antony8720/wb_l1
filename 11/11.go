//Реализовать пересечение двух неупорядоченных множеств.
package main

import (
	"fmt"
)


func intersection(x, y []int) (z []int){
	m := make(map[int]int)
	for _, value := range x{
		m[value]++
	}

	for _, value := range y{
		if m[value] == 1{
			z = append(z,value)
		}
	}
	return z
}

func solution() {
	x := []int{1,3,6,7,8,9,10}
	y := []int{2,3,5,8,10,11}
	z := intersection(x,y)
	fmt.Println(z)
}

func main() {
	solution()
}