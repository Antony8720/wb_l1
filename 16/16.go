package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	l := len(arr)
	if l < 2 {
		return arr
	}
	less := make([]int, 0)
	more := make([]int, 0)
	pivot := arr[0]
	for _, value := range arr[1:] {
		if value > pivot{
			more = append(more, value)
		} else {
			less = append(less, value)
		}
	}
	arr = append(quickSort(less), pivot)
	arr = append(arr, quickSort(more)...)

	return arr
}

func main() {
	arr := []int{10, 4, 1, 6, 2, 5, 8, 3, 9, 7, 0}
	fmt.Println(quickSort(arr))
}