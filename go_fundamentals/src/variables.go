package main

import (
	"fmt"
	"reflect"
)

var (
	name = "Leszek"
	course = "Go fundamentals"
	module = 3.2
)

func main() {
	fmt.Println("Name is", name, "and is of type", reflect.TypeOf(name))
	fmt.Println("Module is", module, "and is of type", reflect.TypeOf(module))
}