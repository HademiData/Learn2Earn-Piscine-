package main

import "fmt"

func CountStr(str string, chr rune) int {
	st := []rune(str)

	counter := 0
	for _, ch := range st {
		if ch == chr {
			counter++
		}
	}
	return counter
}

func main() {
	fmt.Println(CountStr("Hello", ' '))
	fmt.Println(CountStr("GoLaaaang", 'a'))
	fmt.Println(CountStr("Adewale AAAAAAAAAfolabi", 'A'))
	fmt.Println(CountStr("Adewale  7 Afolabi 7", '7'))
}
