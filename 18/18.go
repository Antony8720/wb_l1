// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type counter struct{
	sync.Mutex
	number int
}

func (c *counter) Inc(){
	c.Lock()
	c.number++
	c.Unlock()
}

type atCounter struct{
	number uint64
}

func (atc *atCounter) Inc() {
	atomic.AddUint64(&atc.number, 1)
}

func main() {
	c := counter{}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.number)
	atc := atCounter{}
	for i := 0; i < 1000; i++{
		wg.Add(1)
		go func() {
			defer wg.Done()
			atc.Inc()
		}()
	}
	wg.Wait()
	fmt.Println(c.number)

}