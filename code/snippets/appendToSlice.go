package main

import (
	"fmt"
)

func main() {
	var test = make([]int, 1) // Create an integer slice, of length 1
	fmt.Println(test)         // Printing the array yeilds: [0]
	test = append(test, 1)    // Append the number '1' to the array
	fmt.Println(test)         // Printing the array yeilds: [0 1]
}
