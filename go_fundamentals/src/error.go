package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("heya")

	if err != nil {
		fmt.Println("Error returned was:", err)
	}
}