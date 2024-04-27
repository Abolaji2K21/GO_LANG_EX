package main

import "fmt"

func main() {

	const valueOne = 239.99
	const valueTwo = 129.75
	const valueThree = 99.95
	const valueFour = 350.89
	var totalValue float64

	for {
		fmt.Println("Welcome to the SalesCommission Calculator Created By BeeJay")
		var ItemNumber int
		fmt.Println("Kindly Enter ItemNumber Of Your Choice: ")
		fmt.Scanln(&ItemNumber)
		if ItemNumber == -1 {
			fmt.Print("BYE WE HOPE TO SEE YOU NEXT TIME ")
			break
		} else if ItemNumber < 1 || ItemNumber > 4 {
			fmt.Println("Kindly Enter Another Item Number Ranging From One To Four :")
			continue
		} else {
			if ItemNumber == 1 {
				totalValue += valueOne
			} else if ItemNumber == 2 {
				totalValue += valueTwo
			} else if ItemNumber == 3 {
				totalValue += valueThree
			} else if ItemNumber == 4 {
				totalValue += valueFour
			}

		}

		fmt.Println("Total Value of the Items Selected is :", totalValue)

		var salesPersonEarning = 200 + (0.09 * totalValue)
		fmt.Println("Total sales earning is :", salesPersonEarning)

	}
}
