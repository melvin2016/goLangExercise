// Explanation Of Callback Functions
package main

import "fmt"

func main() {
	num := []int{23, 34, 54, 32, 12}

	// callback is declared as second argument to filter() function.
	filNum := filter(num, func(n int) bool {
		return n > 30
	})

	fmt.Println(filNum)
}

// filter() function which accepts : a slice of int , callback function
func filter(n []int, cb func(int) bool) []int {
	arr := []int{}
	for _, v := range n {
		if cb(v) {
			arr = append(arr, v)
		}
	}
	return arr
}
