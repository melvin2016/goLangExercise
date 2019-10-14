package main

import "fmt"

type specs struct {
	processor string
	ram       int
	battery   int
}

type mobile struct {
	name             string
	yearManufactured int
	warrantyPeriod   float64
	specs
}

func main() {

	// 1 - way of intialising a struct.
	oneplus6t := mobile{
		name:             "Oneplus 6T",
		yearManufactured: 2018,
		warrantyPeriod:   1.5,
		// Embedding Structs
		specs: specs{
			processor: "Snapdragon 845",
			ram:       8,
			battery:   3500,
		},
	}

	// Struct of colors
	colors := []string{"red", "black", "green", "red", "blue"}
	change1stColor(colors)
	fmt.Println(colors)

	// Changing name
	oneplus6t.updateName("1+ 6T")
	// Displaying all key-values
	oneplus6t.print()

}

/*
	Update values in a struct is a pass by value
	so we have to use pointers in go.
*/
func (m *mobile) updateName(name string) {
	(*m).name = name
}

// Print Function attached to type mobile struct
func (m mobile) print() {
	fmt.Printf("\n%+v\n", m)
}

// Structs gets passed by reference and thus no need of pointers
func change1stColor(c []string) {
	c[0] = "yellow"
}
