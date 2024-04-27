package main

import (
	"fmt"
)

func main() {
	var count int
	fmt.Print("How many values do you want to input? ")
	fmt.Scan(&count)

	if count < 2 {
		fmt.Println("Please input at least two values.")
		return
	}

	var num, minimum, maximum, sum int
	fmt.Print("Enter value 1: ")
	fmt.Scan(&num)
	minimum = num
	maximum = num
	sum += num

	for count := 2; count <= count; count++ {
		fmt.Print("Enter value ", count, ": ")
		fmt.Scan(&num)
		sum += num
		if num < minimum {

			minimum = num
		}
		if num > maximum {
			maximum = num
		}
	}

	fmt.Println("Minimum value:", minimum)
	fmt.Println("Maximum value:", maximum)
	fmt.Println("Sum of extremes:", minimum+maximum)
}
