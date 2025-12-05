package main

import (
	"fmt"
	"math"
)

func squre() {
	var val float64
	fmt.Println("Enter the value for finding squre root: ")
	fmt.Scan(&val)

	output := math.Sqrt(val)
	fmt.Println("Squre root is: ", output)
}
