package main

import "fmt"

func main() {

	arr := [5]int{45, 46, 67, 88, 99}

	fmt.Println("Sum :", arr[2]+arr[4])

	i := 0
	j := len(arr) - 1
	slice := arr[i+2 : j]

	fmt.Println("Sum :", slice)
}
