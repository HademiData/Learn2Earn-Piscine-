package main

import (
	"fmt"
)

func ZipString(s string) string {
	if s == "" {
		return ""
	}

	count := 1
	res := ""

	for i := 0; i < len(s); i++ {
		// if next char is same, count++
		if i+1 < len(s) && s[i] == s[i+1] {
			count++
		} else {
			// add current count and character
			res += fmt.Sprintf("%d%c", count, s[i])
			count = 1 // reset for next
		}
	}
	return res
}

func main() {
	fmt.Println(ZipString("YouuungFellllas"))
	fmt.Println(ZipString("Thee quuick browwn fox juumps over the laaazy dog"))
	fmt.Println(ZipString("Helloo Therre!"))
}
