package main

import "fmt"

func main() {

	var hm = map[string]int{
		"Anand":  39,
		"Cannon": 34,
		"Daniel": 25,
		"Erwan":  33,
		"Dirk":   24,
	}

	for key, value := range hm {
		if value == 34 {
			hm[key] = value + 1
		}
		fmt.Printf("Key: %s, Value: %d \n", key, hm[key])
	}
}
