package main

import "fmt"

func main() {
	var accountNumber int
	var beginningBalance int
	var totalItemCharge int
	var totalCredit int
	var creditLimit int

	for {
		fmt.Print("Enter account number (or enter -1 to exit): ")
		fmt.Scanln(&accountNumber)

		if accountNumber == -1 {
			fmt.Println("Thanks Again For your validity check")
			break
		}

		fmt.Println("Awesome So Kindly Enter Your Beginning Balance")
		fmt.Scanln(&beginningBalance)

		fmt.Println("You Are Just One Step Away:")
		fmt.Println("Kindly Enter The Total Item Charge For This Month:")
		fmt.Scanln(&totalItemCharge)

		fmt.Println("At this Stage I know HowTired You Are Getting To put In All Of this Details But we are Close than to the finish line ")
		fmt.Println("Kindly enter the total credit of item for this month")
		fmt.Scanln(&totalCredit)

		fmt.Print("Enter allowed credit limit: ")
		fmt.Scanln(&creditLimit)

		var newBalance = beginningBalance + totalItemCharge - totalCredit

		fmt.Println("Thanks for your patronage Your new balance for accountNumber : ", accountNumber, " is: ", newBalance)

		if newBalance > creditLimit {
			fmt.Println("Credit limit exceeded")
		} else {
			fmt.Println("Awesome You are still within range")
		}

	}

}
