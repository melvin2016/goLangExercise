// UNDESTANDING METHOD SETS

package main

import "fmt"

// Floater Interface
type floater interface {
	float() float64
}

// decimal type with base type int
type decimal int

func main() {

	// QUESTION -
	// Why is compiler restricting
	// the use of values and pointers
	// in reciever function ?

	// ANSWER -
	// Because not every time the values
	// can be accessed using its address
	// and there is a need for specification to
	// avoid these nuances in golang.
	// They are called, the Method sets.

	// EXAMPLE

	// THIS IS VALID -
	// because, its getting an address in the memory
	// by being initialised with a variable.
	d := decimal(3)
	d.PrintDecimal()

	// THIS IS INVALID -
	// because, not getting an address in memory
	// and trying to access the PrintDecimal() function
	// which has a pointer to decimal type.
	// decimal(3).PrintDecimal() /* Uncomment this Line to see the error*/

	// and shows error -
	// cannot call pointer method on decimal(3)
	// cannot take the address of decimal(3)

	// Needs to be passing the pointer
	// of the decimal type because of the
	// methods sets specifications.
	showFloatValue(&d)

}

func (d *decimal) PrintDecimal() {
	fmt.Printf("The Decimal is : %d\n", *d)
}

// pointer type function receivers can only be implemented by pointer values
func (d *decimal) float() float64 {
	return float64(*d)
}

func showFloatValue(f floater) {
	fmt.Println(f.float())
}
