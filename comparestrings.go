package main

import "fmt"

// Write a Go function that takes two strings as arguments.
// The function should return the first string if all of its characters exist anywhere in the second string, regardless of order.
// If any character from the first string is missing in the second string, the function should return an empty string.

func TwoString(val string, val2 string) string {
	str1 := []rune(val)
	str2 := []rune(val2)

	for _, ch1 := range str1 {
		found := false
		for _, ch2 := range str2 {
			if ch1 == ch2 {
				found = true
				break
			}
		}
		if !found {
			return ""
		}
	}
	return val
}

func main() {
	fmt.Println(TwoString("Hello", "Hello World"))
	fmt.Println(TwoString("Golaaang", "Ganagaolawdjf m"))
	fmt.Println(TwoString("zxb", "Ganagaolawdjf m"))
}
