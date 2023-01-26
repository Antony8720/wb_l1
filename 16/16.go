//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.
package main

import (
	"fmt"
	"math/rand"
)

func quickSort1(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}
	less := make([]int, 0)
	more := make([]int, 0)
	pivot := arr[0]
	for _, v := range arr[1:] {
		if v > pivot {
			more = append(more, v)
		} else {
			less = append(less, v)
		}
	}
	arr = append(quickSort1(less), pivot)
	arr = append(arr, quickSort1(more)...)
	return arr
}


func quickSort2(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1

	// Выбираем случайным образом опорный элемент
	pivotIndex := rand.Intn(len(arr))

	// Меняем местами опорный элемент и правый
	arr[pivotIndex], arr[right] = arr[right], arr[pivotIndex]

	// Перемещаем элементы, меньшие опорного, до левого элемента
	for i := range arr {
		if arr[i] < arr[right] {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Ставим опорный элемент после последнего элемента, меньшего чем опорный
	arr[left], arr[right] = arr[right], arr[left]

	// Do it for two parts of current arr
	quickSort2(arr[:left])
	quickSort2(arr[left+1:])

	return arr
}

func main() {
	arr := []int{10, 4, 1, 6, 2, 5, 8, 3, 9, 7, 0}
	sorted1 := quickSort1(arr)
	fmt.Println(sorted1)
	sorted2 := quickSort2(arr)
	fmt.Println(sorted2)
}