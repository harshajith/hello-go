package main

import (
	"fmt"
	"hello/greet"
)

func main() {
	fmt.Println("Variable from nested package: " + greet.Morning)
	fmt.Println("Package variable usage: ", snapshot)
	fmt.Println("Constant Usage: ", greet.TestConst)

	fmt.Println(getName(" Halgaswatta"))

	stringOps()
	arrayOps("", "")

	mapOps()
	pointerOps()

	structOps()

}
