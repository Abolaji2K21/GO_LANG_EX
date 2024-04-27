package main

import (
	"fmt"
)

func main() {
	var counter int = 1
	var largest int = 0

	for counter <= 10 {
		fmt.Printf("Enter number %d: ", counter)
		var number int
		fmt.Scanln(&number)

		if number > largest {
			largest = number
		}

		counter++
	}

	fmt.Println("The largest number is:", largest)
}
