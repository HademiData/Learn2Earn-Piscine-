// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func quadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 && col == 0) || (row == y-1 && col == x-1) {
				fmt.Print("/")

			} else if (row == 0 && col == x-1) || (row == y-1 && col == 0) {
				fmt.Print("\\")

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
	quadA(5, 4)
}
