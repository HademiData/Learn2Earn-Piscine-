// Goal: Print all possible combinations of two different two-digit numbers in ascending order.
// Format: Combinations separated by a comma and a space, e.g., 00 01, 00 02, ..., 98 99.
// Expected Function: func PrintComb2()

package main

import "fmt"

func main() {

	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				for l := 0; l <= 9; l++ {
					fmt.Printf("%d%d", i, j)
					fmt.Printf(" ")
					fmt.Printf("%d%d", k, l)

					if !(i == 9 && j == 8 && k == 9 && l == 9) {
						fmt.Println(", ")
					}
				}

			}

		}
	}
	fmt.Print("\n")
}
