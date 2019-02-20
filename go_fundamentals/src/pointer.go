package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := "elo"

	ptr := &a

	fmt.Println("Value of A:", *ptr)
	fmt.Println("Address of A:", ptr)
	fmt.Println("Type of A address:", reflect.TypeOf(ptr))
}