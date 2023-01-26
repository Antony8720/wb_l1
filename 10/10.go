//Дана последовательность температурных колебаний: -25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
//Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
package main

import "fmt"

func solution() {
	sl := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	temp := make(map[int][]float64)
	for _, value := range sl{
		i := int(value) - int(value) % 10
		temp[i] = append(temp[i], value)
	}
	fmt.Println(temp)
}

func main() {
	solution()
}