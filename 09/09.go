package main

import "fmt"

func reader(sl []int, ch1 *chan int) {
	for _, value := range sl{
		*ch1 <- value
	}
	close(*ch1)
}

func writer(ch1 *chan int, ch2 *chan int) {
	for value := range *ch1{
		*ch2 <- value * 2
	}
	close(*ch2)
}

func solution() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	sl := []int{2,4,6,8,10}
	go reader(sl, &ch1)
	go writer(&ch1, &ch2)

	for value := range ch2{
		fmt.Println(value)
	}


}

func main() {
	solution()
}