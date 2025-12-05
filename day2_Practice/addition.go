package main

import "fmt"

func add() {
	var num1 int
	var num2 int
	fmt.Println("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Println("Enter the second number: ")
	fmt.Scan(&num2)

	fmt.Println("Addtion is: ", num1+num2)
}
