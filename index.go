package main

import "fmt"

func Index(s string, toFind string) int {

	runesS := []rune(s)
	runesFind := []rune(toFind)

	if len(toFind) <= 0 {
		return 0
	}

	for i := 0; i <= len(s)-len(toFind); i++ {
		match := true
		for j := 0; j < len(toFind); j++ {
			if runesS[i+j] != runesFind[j] {
				match = false
				break
			}
		}
		if match {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Println(Index("adewale", "wale"))
}
