// Explanation On defer keyword
package main

import "fmt"

func main() {
	// defer keyword is used to postpone a function
	// call to the exit of its lexical function scope.
	// so that world() function will be called at exit of main function.
	// and hello() function will be called first.

	/* -------------------------- */

	// Useful for readability of code
	// most commonly used with file closing.
	// eg :
	// file.open()
	// defer file.close() -- useful for code readability.
	defer world()
	hello()
}

func hello() {
	fmt.Println("Hello")
}

func world() {
	fmt.Println("world!")
}
