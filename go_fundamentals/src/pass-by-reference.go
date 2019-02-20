package main

import (
	"fmt"
)

func main() {
	name := "Leszek"
	course := "Go fundamentals"

	fmt.Println("\nHi", name, "you are watching", course)

	changeCourse(&course)

	fmt.Println("\nYou are now watching course", course)
}

func changeCourse(course *string) string {

	*course = "How to shit in the woods"

	fmt.Println("Trying to change course to", course)

	return course
}