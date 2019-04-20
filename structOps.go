package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	salary    Salary
	fullTime  bool
}

type Salary struct {
	amount    int
	insurance bool
}

func structOps() {
	var hala Employee

	fmt.Println("employee struct:", hala)

	hala.firstName = "Harsha"
	hala.lastName = "Halgaswatta"
	hala.salary = Salary{32, false}
	hala.fullTime = false

	var singee = Employee{
		firstName: "Buddini",
		lastName:  "Malwatta",
		salary:    Salary{54, true},
		fullTime:  true,
	}

	ross := Employee{"Ross", "Geller", Salary{54, false}, true}

	fmt.Println("employee struct", hala)

	fmt.Println("employee struct", singee)

	fmt.Println("employee struct", ross)

	savein := &Employee{
		firstName: "Savein",
		lastName:  "Hiranya",
		salary:    Salary{32, false},
		fullTime:  true,
	}

	savein.firstName = "SSS"
	fmt.Println("employee struct", savein)

	changeName(ross)
	fmt.Println("Ross firstname:", ross.firstName)

	changeFirstName(savein)
	fmt.Println("New firstname: ", savein.firstName)
}

func changeName(e Employee) {
	e.firstName = "Sethi"
	fmt.Println("Changed Name:", e.firstName)
}

func changeFirstName(e *Employee) {
	e.firstName = "Pramudith"
}
