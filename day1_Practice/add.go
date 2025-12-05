package main

import "fmt"

func add() {
	var num1, num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println("sum is: ", num1+num2)
}
