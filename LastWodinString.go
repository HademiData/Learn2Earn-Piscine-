package main

import "fmt"

func isLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')
}

func LastWord(s string) string {
	str := []rune(s)

	if len(str) == 0 {
		return ""
	}

	end := len(str) - 1
	for end >= 0 && (str[end] == ' ' || !isLetter(str[end])) {
		end--
	}

	if end < 0 {
		return ""
	}

	start := end
	for start >= 0 && isLetter(str[start]) {
		start--
	}

	return string(str[start+1 : end+1])

}

func main() {
	fmt.Println(LastWord("Hello"))                                             // Hello
	fmt.Println(LastWord("GoLaaaang"))                                         // GoLaaaang
	fmt.Println(LastWord("Adewale AAAAAAAAAfolabi"))                           // AAAAAAAAAfolabi
	fmt.Println(LastWord("Adewale is a 7 Afolabi Ade the first of his name!")) // name
	fmt.Println(LastWord("ade food yummy"))                                    // yummy
	fmt.Println(LastWord("wale paul pip !"))                                   // pip
	fmt.Println(LastWord("    "))                                              // (nothing)
	fmt.Println(LastWord("  hello  "))                                         // hello
}
