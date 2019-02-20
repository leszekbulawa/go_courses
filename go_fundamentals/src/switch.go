package main

import (
	"fmt"
)

func main() {
	switch "docker" {
	case "rkt" :
		fmt.Println("Recommended rkt courses...")
	case "docker":
		fmt.Println("Recommended docker courses...")
	case "cri-o":
		fmt.Println("Reccomended cri-o courses...")
	default:
		fmt.Println("No matches.")
	}
}