package main

import "fmt"

func pointerOps() {
	var a *int

	b := 54
	a = &b

	fmt.Printf("pointer type: %T of address %v, value : %v", a, a, *a)

	*a = 43
	fmt.Printf("new value %v", *a)

	x := 32
	changeValue(&x)

	fmt.Println("x value after change is: v", x)

	var zz = [3]int{3, 5, 6}
	changeArray(&zz)

	fmt.Println("value is:", zz[2])
}

func getAge(a *int) int {
	return *a + 32
}

func changeValue(a *int) {
	*a = 67
}

func changeArray(arr *[3]int) {
	arr[2] = 100
}
