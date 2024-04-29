package main

import (
	"fmt"
)

func main() {
	first := 10
	second := 3
	sum := sumTwoNumbers(first, second)
	sumOne := sumTwoNumberAgain(first, second)
	fmt.Println("Sum:", sum)
	fmt.Println("Sum:", sumOne)

}

func sumTwoNumbers(first, second int) int {
	for second != 0 {
		first++
		second--
	}
	return first
}

func sumTwoNumberAgain(first, second int) int {
	return first - (-second)
}
