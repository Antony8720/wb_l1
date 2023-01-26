//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.
package main

import "fmt"

func solution() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}

	m := make(map[string]bool)

	for _, value := range sl{
		m[value] = true
	}

	fmt.Println(m)

}

func main() {
	solution()
}