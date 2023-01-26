package main

import (
	"fmt"
)

func partition(arr []int) int{
	pivot := arr[len(arr)/2]

	left := 0
	right := len(arr) - 1

	for{
	for; arr[left] < pivot; left++{
	}
	for; arr[right] > pivot; right--{
	} 

	if left >= right {
		return right
	}
	arr[left], arr[right] = arr[right], arr[left]
	}
}

func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	split := partition(arr)
	quickSort(arr[:split])
	quickSort(arr[split:])
}

func main() {
	arr := []int{10, 4, 1, 6, 2, 5, 8, 3, 9, 7, 0}
	quickSort(arr)
	fmt.Println(arr)
}