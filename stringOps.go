package main

import "fmt"

var snapshot = "1.0.0-SNAPSHOT"

func getName(s string) string {
	return "My name is Harsha" + s
}

func stringOps() {
	var s = "Hello World"
	fmt.Printf(s)

	for a := 0; a < len(s); a++ {
		fmt.Printf(" char is %c", s[a])
	}

}
