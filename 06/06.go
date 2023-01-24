//Реализовать все возможные способы остановки выполнения горутины.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Способ 1
// Остановка горутины с помощью контекста с таймаутом/дедлайном
func solution1() {
	ch := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), time.Second*3)

	go func() {
		for {
			ch <- time.Now().Minute()
			time.Sleep(time.Second)
		}
	}()

	for {
		select {
		case v := <-ch:
			fmt.Println(v)
		case <-ctx.Done():
			return
		}
	}
}

// Способ 2
// Остановка горутины с помощью сигнального канала
func solution2() {
	ch := make(chan bool)
	go func() {
		for {
			fmt.Println("Ожидание закрытия")
			select {
			case <-ch:
				return
			}
		}
	}()
	time.Sleep(time.Second * 3)
	ch <- true
	fmt.Println("Горутина остановлена")

}

// Способ 3 
// Остановка с помощью WaitGroup
func solution3() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func (){
		defer wg.Done()
		fmt.Println("горутина запущена")
		time.Sleep(time.Second * 2)
	}()
	wg.Wait()
	fmt.Println("Горутина остановлена")

}

// Способ 4 
// Остановка с помощью context.WithCancel
func solution4() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			fmt.Println("Ожидание закрытия")
			select {
			case <-ctx.Done():
				return
			}
		}
	}()
	time.Sleep(time.Second * 2)
	cancel()
	fmt.Println("Горутина остановлена")

}

// Способ 5 
// Остановка с помощью закрытия канала
func solution5() {
	ch := make(chan bool)
	go func() {
		for {
			fmt.Println("Ожидание закрытия")
			_, ok := <- ch
			if !ok {
				return
			}
			}
		}()
	time.Sleep(time.Second * 3)
	close(ch)
	fmt.Println("Горутина остановлена")

}

func main() {
	// solution1()
	// solution2()
	// solution3()
	// solution4()
	solution5()
}
