package main

import (
	"fmt"
)

func main() {
	myCourses := make([]string, 5, 10)

	fmt.Printf("Length is: %d\nCapacity is: %d", len(myCourses), cap(myCourses))
}