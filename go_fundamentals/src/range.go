package main

import (
	"fmt"
)

func main() {
	coursesInProg := []string{"Docker Deep Dive", "Docker Clustering", "Go fundamentals"}
	coursesCompleted := []string{"How to shit in the woods", "Docker Deep Dive", "Python fundamentals"}

	for _, i := range coursesInProg {
		for _, j := range coursesCompleted{
			if i == j {
				fmt.Println("Hey, there's a clash! Course", j, "is in both lists!")
			}
		}
	}
}