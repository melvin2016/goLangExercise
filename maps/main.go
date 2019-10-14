// Contains my testings based on map type in Golang.
package main

import "fmt"

func main() {
	/*
		1. map type in Go is a reference type.
		2. it can only contain key of same type and
		   values of same type.
	*/

	// 1st way of initailising and declaring a map type.
	carsToColors := map[string]string{
		"audi":    "black",
		"bmw":     "red",
		"bugatti": "yellow",
	}
	// manipulating contents declared while initialising using square brackets [].
	carsToColors["audi"] = "red"

	// 2nd way of initialising map type.
	var nameToPass = make(map[string]string)
	// declaring key-values in a map needs to have square brackets [].
	nameToPass["melvin"] = "sach345"
	nameToPass["basilroy"] = "basil4wayne123"

	// 3rd way of initialising a map type using make() function.
	intToString := make(map[int]string)
	// manipulating contents inside of intTostring map.
	intToString[1] = "One"
	intToString[2] = "Two"

	// deleting an entry in map using delete() function.
	delete(intToString, 1)
	delete(nameToPass, "melvin")

	// Printing contents inside maps
	fmt.Println(carsToColors)
	fmt.Println(intToString)
	fmt.Println(nameToPass)
}
