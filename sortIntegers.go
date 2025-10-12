// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func SortInt(arr []int) {

	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Print(arr)

}

func main() {
	SortInt([]int{5, 2, 1, 2, 3, 5, 33, 100, 7})
}
