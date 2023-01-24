// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.
// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(ctx context.Context, i int, ch <-chan int, wg *sync.WaitGroup){
	defer wg.Done()
	for {
		select {
		case num:= <- ch:
			fmt.Printf("Воркер № %v прочитал значение %v\n", i, num)
		case <- ctx.Done():
			fmt.Printf("Воркер № %v закрыл поток\n", i)
			return

		}
	}
}
func solution() {
	wg := sync.WaitGroup{}
	nWorkers := 0
	ch := make(chan int)
	sChan := make(chan os.Signal ,1)
	signal.Notify(sChan, syscall.SIGINT)
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Введите количество воркеров")
	fmt.Scan(&nWorkers)

	for i := 0; i < nWorkers; i++ {
		wg.Add(1)
		go worker(ctx, i, ch, &wg)
	}

	for {
		select{
		case <- sChan:
			cancel()
			wg.Wait()
			fmt.Println("Программа остановлена")
			return
		default:
			ch <- rand.Intn(1000)
			time.Sleep(time.Millisecond *500)

		}

	}
}

func main() {
	solution()
}