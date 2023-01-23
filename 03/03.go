// Дана последовательность чисел: 2,4,6,8,10.
// Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var array = []int64{2,4,6,8,10}
var mu sync.Mutex

// 1 решение
// Используем Wait Group.
// Перед выполнением горутины увеличивавем счетчик на 1.
// После выполнения уменьшаем счетчик на 1.
// В конце функции ожидаем завершения всех горутин.
// Для избежания race condition используем мьютекс
func firstSolution() {
	var sum int64
	wg := sync.WaitGroup{}
	for _, v := range array{
		wg.Add(1)
		go func(v int64){
			mu.Lock()
			sum += v * v
			mu.Unlock()
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println(sum)
}

// 2 решение
// Используем Wait Group.
// Перед выполнением горутины увеличивавем счетчик на 1.
// После выполнения уменьшаем счетчик на 1.
// В конце функции ожидаем завершения всех горутин.
// Для избежания race condition используем atomic
func secondSolution() {
	var sum int64
	wg := sync.WaitGroup{}
	for _, v := range array{
		wg.Add(1)
		go func(v int64){
			atomic.AddInt64(&sum,v * v)
			wg.Done()
		}(v)
	}
	wg.Wait()
	fmt.Println(sum)
}

// 3 решение
// Используем буферизированный канал.
// В горутине записываем в буфер канала вычисленный квадрат числа
// Забираем из буфера канала вычисленное значение и выводим его в консоль
func thirdSolution() {
	var sum int64
	ch := make(chan int64,5)
	for _, v := range array{
		go func(v int64){
			ch <- v * v
		}(v)
		sum += <- ch
	}
	fmt.Println(sum)
}


// 4 решение
// Используем небуферизированный канал.
// В горутине записываем в канал вычисленный квадрат числа
// Забираем из канала вычисленное значение и выводим его в консоль
func fourthSolution() {
	var sum int64
	ch := make(chan int64)
	for _, v := range array{
		go func(v int64){
			ch <- v * v
		}(v)
		sum += <- ch
	}
	fmt.Println(sum)
}



func main() {
	firstSolution()
	secondSolution()
	thirdSolution()
	fourthSolution()
}