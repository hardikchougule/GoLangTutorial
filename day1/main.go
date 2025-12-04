package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Println("Enter your age: ")
	fmt.Scan(&age)

	fmt.Println("Your name is ", name, "and age is ", age)
}

// for compilation command is:- go mod init go-lang-totrial
// for execution command is :- go run main.go
