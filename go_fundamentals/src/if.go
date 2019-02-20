package main

import (
	"fmt"
)

func main() {
	firstRank := "128"
	secondRank := "256"

	if firstRank < secondRank {
		fmt.Println("\nSecond rank is bigger.")
	} else if firstRank > secondRank {
		fmt.Println("\nFirst rank is bigger.")
	} else {
		fmt.Println("\nRanks are equal.")
	}
}