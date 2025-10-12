package main

import "fmt"

// Write a Go function that returns the first word of a sentence.
// The function takes in a sentence as an argeument.

func FindFirstWrd(val string) string {
	str := []rune(val)

	if len(str) == 0 {
		return ""
	}

	start := 0
	for start >= 0 && (str[start] >= 'a' && str[start] <= 'z') || (str[start] >= 'A' && str[start] <= 'Z') {
		start++
	}

	if start < 0 {
		return ""
	}

	return string(str[0 : start+1])

}

func main() {
	fmt.Println(FindFirstWrd("Hello Adewale"))
	fmt.Println(FindFirstWrd("Golang is really useful"))
	fmt.Println(FindFirstWrd("mango is a fruit."))
}
