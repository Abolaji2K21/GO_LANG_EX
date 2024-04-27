package main

import "fmt"

func main() {
	fmt.Println("Sum of numbers ranging from one to hundred")

	var sum int

	for count := 1; count <= 100; count++ {
		sum += count
		fmt.Println("Count:", count, "While Sum: ", sum)
	}
}
