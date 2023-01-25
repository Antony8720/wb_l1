package main

import (
	"fmt"
)

func solution(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("int: ", v)
	case string:
		fmt.Println("string: ", v)
	case bool:
		fmt.Println("bool: ", v)
	case chan interface{}:
		fmt.Println("chan: ", v)
	default:
		fmt.Println("Unknown type: ", v)
		
	}
	
}

func main() {
	channel := make(chan interface{})
	solution(0)
	solution("str")
	solution(true)
	solution(channel)
	solution(0.1)
}