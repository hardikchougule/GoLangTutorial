package main

import "fmt"

func sub() {
	var num1 int
	var num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println("Substration is: ", num1-num2)
}
