package main

import "fmt"

func revstr(str string) string {
	st := []rune(str)

	length := len(st)
	for i := 0; i <= length/2; i++ {
		j := length - i - 1
		st[i], st[j] = st[j], st[i]
	}
	return string(st)

}

func main() {
	fmt.Println(revstr("Hello"))
	fmt.Println(revstr("GoLang"))
	fmt.Println(revstr("Afolabi"))
}
