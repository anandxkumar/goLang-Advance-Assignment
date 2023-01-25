package main

import "fmt"

func main() {

	var f float64
	f = 4.5

	adr := &f

	fmt.Println("Address of f: ", adr)
	fmt.Println("Initial Value of f: ", *adr)

	*adr = 2.1

	fmt.Println("Updated value of f: ", f)
}
