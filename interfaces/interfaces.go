// Interface types are a contract to make code better and manageable.
package main

import "fmt"

// house interface giving out
// specific information of
// getFloorNumber Function.
type house interface {
	getFloorNumber() int
}

// struct flat which implements the house interface,
// because it defines getFloorNumber() function.
type flat struct{}

// struct bunker which implements the house interface,
// because it defines getFloorNumber() function.
type bunker struct{}

func main() {
	ft := flat{}
	bk := bunker{}
	printFloorNumber(ft)
	printFloorNumber(bk)
}

// function of flat struct implementing the
// getFloorNumber() function of house interface.
func (flat) getFloorNumber() int {
	return 32
}

// function of bunker struct implementing the
// getFloorNumber() function of house interface.
func (bunker) getFloorNumber() int {
	return 0
}

// Interface house helps in not writing
// 2 different function for 2 different
// structs since both use house interface as base.
func printFloorNumber(h house) {
	fmt.Println("The floor number is: ", h.getFloorNumber())
}
