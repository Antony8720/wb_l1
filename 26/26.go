// Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
// Функция проверки должна быть регистронезависимой.
package main

import (
	"fmt"
	"strings"
)

func solution1(str string) bool{
	for i := 0; i < len(str); i++{
		for j := i+1; j < len(str); j++{
			if str[i] == str[j]{
				return false
			}
		}
	}
	return true
}

func solution2(str string) bool{
	strLow := strings.ToLower(str)
	runes := []rune(strLow)
	m := make(map[rune]int)

	for _, value := range runes {
		if m[value] == 1{
			return false
		} else {
			m[value]++
		}
	}
	return true
}

func main() {
	fmt.Println(solution1("abcd"))
	fmt.Println(solution1("abCdefAa"))
	fmt.Println(solution1("aabcd"))

	fmt.Println(solution1("abcd"))
	fmt.Println(solution2("abCdefAa"))
	fmt.Println(solution2("aabcd"))
}