//Разработать программу, которая будет последовательно отправлять значения в канал,
// а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func reader(ch *chan int){
	for v := range *ch {
		fmt.Printf("Получено значение %v\n", v)
	}
}

func writer(ch *chan int) {
	for {
	value := rand.Intn(1000)
	*ch <- value
	fmt.Printf("Отправлено значение %v\n",value)
	}
}

func solution1() {
	var n int
	fmt.Println("Введите время работы программы в секундах")
	fmt.Scan(&n)
	seconds := n * int(time.Second)
	ch := make(chan int)
	go writer(&ch)
	go reader(&ch)
	time.Sleep(time.Duration(seconds))
	close(ch)

}

func solution2() {
	const n = 1
	ch := make(chan int)
	ctx, _ := context.WithTimeout(context.Background(), time.Second * n)
	go func() {
		for{
		value := rand.Intn(1000)
		ch <- value
		fmt.Printf("Отправлено значение %v\n",value)
		time.Sleep(time.Second)
		}
	}()

	for {
		select{
		case v:= <- ch:
			fmt.Printf("Получено значение %v\n", v)
		case <- ctx.Done():
			return
		}
	}
}

func main() {
	// solution1()
	solution2()
}