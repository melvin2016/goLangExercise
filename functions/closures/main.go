// Creation Of Closures And Use Of Anonymous Functions.
package main

import "fmt"

func main() {
	// The incrementWrapper() return a function with x value enclosed in them
	// This is a closure pattern.
	// Closure pattern is used to make variables and function private.
	increment := incrementWrapper()

	// increment variable is declared with a function and can be called.
	fmt.Println("Value: ", increment())
	fmt.Println("Value: ", increment())
	fmt.Println("Value: ", increment())
	fmt.Println("Value: ", increment())

}

func incrementWrapper() func() int {
	x := 0
	// this return an anonymous function
	return func() int {
		x++
		return x
	}
}
