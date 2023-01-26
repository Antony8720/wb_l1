package main

import (
	"fmt"
)

func binarySearch(arr []int, elem int) bool{
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		if arr[mid] == elem {
			return true
		}
		if arr[mid] < elem{
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if low == len(arr) || arr[low] != elem {
		return false
	}
	return true	
}

func main() {
	arr := []int{1,2,5,15,27,45,78,101}
	fmt.Println(binarySearch(arr, 45))
}