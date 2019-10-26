// Explanation Of Variadic Functions {arguments and parameters}
package main

import "fmt"

func main() {
	data := []string{"Melvin", "Nibin", "Basil", "Raneesh", "Akash"}
	showStrings(data...)
}

// Variadic Parameter of str
func showStrings(str ...string) {
	// It eventualy evalutes to a slice of strings
	for i, v := range str {
		fmt.Println(i, "string is : ", v)
	}
}
