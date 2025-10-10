// Goal: Print all possible combinations of two different two-digit numbers in ascending order.
// Format: Combinations separated by a comma and a space, e.g., 00 01, 00 02, ..., 98 99.
// Expected Function: func PrintComb2()
// Click here and start typing.
package main

import "fmt"

func main() {

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					val1 := ((i * 10) + j)
					val2 := ((k * 10) + l)
					if val1 < val2 {
						// Print with leading zeros and proper spacing
						fmt.Printf("%02d %02d", val1, val2)

						// Avoid trailing comma
						if !(val1 == 98 && val2 == 99) {
							fmt.Print(", ")
						}
					}
				}
			}

		}
	}
	fmt.Print("\n")
}
