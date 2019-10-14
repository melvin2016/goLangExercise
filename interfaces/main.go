package main

import "fmt"

type house interface {
	getFloorNumber() int
}
type flat struct{}
type bunker struct{}

func main() {
	ft := flat{}
	bk := bunker{}
	printFloorNumber(ft)
	printFloorNumber(bk)
}

func (flat) getFloorNumber() int {
	return 32
}

func (bunker) getFloorNumber() int {
	return 0
}

func printFloorNumber(h house) {
	fmt.Println("The floor number is: ", h.getFloorNumber())
}
