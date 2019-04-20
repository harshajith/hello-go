package main

import "fmt"

func mapOps() {
	var students = make(map[string]string)

	students["kamal"] = "32"
	students["harsha"] = "34"

	fmt.Println(len(students))
	fmt.Println("map value", students["kamal"])

	harshaAge, harshaOk := students["harsha"]
	fmt.Println(harshaOk, harshaAge)

	for key, value := range students {
		fmt.Println(key, value)
	}

}
