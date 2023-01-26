package main

import (
	"strings"
	"fmt"
)

func solution(s string) {
	words := strings.Fields(s)
	var result string

	for i := len(words) - 1; i >= 0; i--{
		result += words[i] + " "
	}
	fmt.Println(result)
	
}

func main() {
	str := "snow dog sun"
	solution(str)
}