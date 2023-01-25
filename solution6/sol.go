package main

import "fmt"

func main() {

	var a = 67.88
	num1 := int(a)
	fmt.Println(int(num1))

	num2 := 21.5

	var sum float64

	sum = float64(num1) + num2

	fmt.Println("Sum: ", sum)
}
