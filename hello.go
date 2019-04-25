package main

import (
	"fmt"
	"hello-go/greet"
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

	student := Student{
		Firstname: "harsha",
		Lastname:  "halgaswatta",
		age:       32,
	}
	fmt.Println("Fullname:", getFullName(student.Firstname, student.Lastname))
	fmt.Println("Student:", student)
	fmt.Println("Method value:", student.getFullNameMethod())
	fmt.Println("Value not changed :", student.age)
	fmt.Println("pointer reciever method", student.getFullName())
	fmt.Println("Value changed :", student.age)

	testInterfaces()

	readFile()

}

func testInterfaces() {
	var c Department = CSE{description: "computer science and engineering"}
	fmt.Println("Dean is: ", c.dean())

	a := c.(Lab) // Change dynamic type to another implementing static type to invoke its methods
	fmt.Println(a.totalFacilities())
}
