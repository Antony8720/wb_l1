package main

import "fmt"

func solution(elem int, slice []int) []int {
	slice = append(slice[:elem],slice[elem+1:]... )
	return slice
}

func main() {
	slice := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(solution(3, slice))
}