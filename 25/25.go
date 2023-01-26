//Реализовать собственную функцию sleep.
package main

import (
	"fmt"
	"time"	
)

func sleep(d time.Duration) {
    <-time.After(d)
}
func main() {
	fmt.Println("start main")
	d := time.Second * 3
	sleep(d)
	fmt.Println("stop main")
}