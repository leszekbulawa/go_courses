package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Current user:", os.Getenv("USER"))
}