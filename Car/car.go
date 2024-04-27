package main

import "fmt"

func main() {

	var totalMiles float64
	var totalGallons float64
	var tripCount int

	for {
		fmt.Print("Enter miles driven for trip ", tripCount+1, " (or enter -1 to exit): ")
		var miles int
		fmt.Scanln(&miles)

		if miles == -1 {
			break
		}

		fmt.Print("Enter gallons used for trip ", tripCount+1, ": ")
		var gallons int
		fmt.Scanln(&gallons)

		mpg := float64(miles) / float64(gallons)

		fmt.Println("Miles per gallon for trip", tripCount+1, ":", mpg)

		totalMiles += float64(miles)
		totalGallons += float64(gallons)

		tripCount++
	}

	combinedMPG := float64(totalMiles) / float64(totalGallons)

	fmt.Println("Combined miles per gallon for all trips:", combinedMPG)
}
