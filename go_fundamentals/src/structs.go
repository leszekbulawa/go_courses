package main

import (
	"fmt"
)

func main() {

	type courseMeta struct {
		Author string
		Level string
		Rating float64
	}

	// var DockerDeepDive courseMeta
	// DockerDeepDive := new(courseMeta) - pointer

	SelfDefence := courseMeta{
		Author: "Arnold Schwarzenegger",
		Level: "Rookie",
		Rating: 5,
	}

	fmt.Println("Author:", SelfDefence.Author)
	fmt.Println("Level:", SelfDefence.Level)
}