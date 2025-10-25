package main

import "fmt"

func quadC(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if row == 0 && col == 0 {
				fmt.Print("/")
			} else if row == 0 && col == x-1 {
				fmt.Print("\\")
			} else if row == y-1 && col == 0 {
				fmt.Print("\\")
			} else if row == y-1 && col == x-1 {
				fmt.Print("/")
			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func main() {
	quadC(1, 3)
}
