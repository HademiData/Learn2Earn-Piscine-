// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func quadD(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}

	for row := 0; row < y; row++ {
		for col := 0; col < x; col++ {
			if (row == 0 && col == 0) || (col == x-1 && row == 0) {
				fmt.Print("A")

			} else if row == y-1 && (col == x-1 || col == 0) {
				fmt.Print("C")

			} else if row == 0 || row == y-1 || col == 0 || col == x-1 {
				fmt.Print("B")

			} else {
				fmt.Print(" ")

			}
		}
		fmt.Print("\n")

	}
}

func main() {
	quadD(0, 0)
}
