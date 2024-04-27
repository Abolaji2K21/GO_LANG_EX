package main

import (
	"fmt"
)

func main() {
	var largest1, largest2, number int

	fmt.Println("Enter number 1: ")
	fmt.Scanln(&largest1)

	fmt.Println("Enter number 2: ")
	fmt.Scanln(&number)

	if number > largest1 {
		largest2 = largest1
		largest1 = number
	} else {
		largest2 = number
	}

	for i := 3; i <= 10; i++ {
		fmt.Println("Enter number", i, ": ")
		fmt.Scanln(&number)

		if number > largest1 {
			largest2 = largest1
			largest1 = number
		} else if number > largest2 {
			largest2 = number
		}
	}

	fmt.Println("The largest number is:", largest1)
	fmt.Println("The second largest number is:", largest2)
}
