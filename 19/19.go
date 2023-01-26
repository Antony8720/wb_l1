package main

import (
	"fmt"
)


func reverseString1(s string) (res string) {
	for _, value := range s {
		res = string(value) + res
	}
	return res
	
}

func reverseString2(s string) {
	for _, value := range s {
		defer fmt.Print(string(value))
	}
	
}

func reverseString3(s string) {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
	
}

func main() {
	str := "главрыба"
	fmt.Println(reverseString1(str))
	reverseString2(str)
	reverseString3(str)
}