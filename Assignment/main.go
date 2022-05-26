package main

import "fmt"

func main() {
	for i, _ := range [11]int{} {
		result := ""

		if i%2 == 0 {
			result = "even"
		} else {
			result = "odd"
		}

		fmt.Println(fmt.Sprintf("%v", i) + " is " + result)
	}
}
