package main

// Goal: Print all unique combinations of three different digits (a < b < c) in ascending order.
// Format: Combinations separated by a comma and a space, e.g., 012, 013, ..., 789.
// Expected Function: func PrintComb()

import "fmt"

func main() {
	for i := 0; i <= 9; i++ {
		for j := 0; j <= 9; j++ {
			for k := 0; k <= 9; k++ {
				if i < j && j < k {
					fmt.Printf("%d%d%d", i, j, k)

					// print comma and space unless it's the last combo
					if !(i == 7 && j == 8 && k == 9) {
						fmt.Print(", ")
					}
				}
			}
		}
	}
	fmt.Print("\n")
}
