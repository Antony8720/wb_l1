//Реализовать конкурентную запись данных в map.
package main

import (
	"fmt"
	"sync"
)

type ConcurentMap struct{
	sync.RWMutex
	data map[string]int
}

func(cm *ConcurentMap) Get(key string) (int){
	cm.RLock()
	value:= cm.data[key]
	cm.RUnlock()
	return value
}

func(cm *ConcurentMap) Inc(key string) {
	cm.Lock()
	cm.data[key]++
	cm.Unlock()
}

func solution() {
	cm := ConcurentMap{data: make(map[string]int)}
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++{
		wg.Add(1)
		go func ()  {
			defer wg.Done()
			cm.Inc("counter")
		}()
	}
	wg.Wait()
	fmt.Println(cm.Get("counter"))
}

func main() {
	solution()
}