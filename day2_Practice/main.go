package main

import (
	"fmt"
)

func main() {

	//while loop demo
	// i := 1
	// for i <= 5 {
	// 	fmt.Printf("*")
	// }

	//for loop demo
	// for i := 0; i <= 9; i++ {
	// 	for j := 0; j <= i; j++ {
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// }

	//infinitefor
	// for {

	// }

	for {
		var choice int
		fmt.Println("Choose the operation you want to perform: ")
		fmt.Println("1 Addition \n 2 substration \n 3 multiplication \n 4 division \n 5 absulute \n 6 squreroot \n 0 exit")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("This is a Additon Opration Enter 2 Numbers: ")
			add()
		case 2:
			fmt.Println("This is a substration Opration Enter 2 Numbers: ")
			sub()
		case 3:
			fmt.Println("This is a multiplication Opration Enter 2 Numbers: ")
			mul()
		case 4:
			fmt.Println("This is a divison operation Enter 2 Numbers: ")
			div()

		case 5:
			abs()

		case 6:
			squre()
		}
		if choice == 0 {
			fmt.Println("Opration Ended")
			break
		}
	}
}
