package main

import (
	"fmt"
	"math"
)

func abs() {
	var ab float64
	fmt.Print("Enter the number for absulute value: ")
	fmt.Scan(&ab)

	output := math.Abs(ab)
	fmt.Println("Absulute value is: ", output)
}
