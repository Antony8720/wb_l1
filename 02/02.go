//Написать программу, которая конкурентно рассчитает значение квадратов чисел
// взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
	"time"
)

var array = []int{2,4,6,8,10}

// 1 решение
// Используем Wait Group.
// Перед выполнением горутины увеличивавем счетчик на 1.
// После выполнения уменьшаем счетчик на 1.
// В конце функции ожидаем завершения всех горутин
func firstSolution() {
	wg := sync.WaitGroup{}
	for _, v := range array{
		wg.Add(1)
		go func(v int){
			fmt.Println(v*v)
			wg.Done()
		}(v)
	}
	wg.Wait()
}

// 2 решение
// Используем небуферизированный канал.
// В горутине записываем в канал вычисленный квадрат числа
// Забираем из канала вычисленное значение и выводим его в консоль
func secondSolution() {
	ch := make(chan int)
	for _, v := range array{
		go func(v int){
			ch <- v * v
		}(v)
		fmt.Println(<-ch)
	}
}

// 3 решение
// Используем буферизированный канал.
// В горутине записываем в буфер канала вычисленный квадрат числа
// Забираем из буфера канала вычисленное значение и выводим его в консоль
func thirdSolution() {
	ch := make(chan int,5)
	for _, v := range array{
		go func(v int){
			ch <- v * v
		}(v)
		fmt.Println(<-ch)
	}
}

// 4 решение
// Используем time.Sleep для ожидания выполнения всех горутин
func fourthSolution() {
	for _, v := range array{
		go func(v int){
			fmt.Println(v*v)
		}(v)
	}
	time.Sleep(1* time.Second)
}



func main() {
	firstSolution()
	secondSolution()
	thirdSolution()
	fourthSolution()
}