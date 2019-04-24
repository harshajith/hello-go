package main

type Student struct {
	Firstname string
	Lastname  string
	age       int
}

func getFullName(firstName string, lastName string) string {
	return firstName + " " + lastName
}

func (e Student) getFullNameMethod() string {
	e.age = 43
	return e.Firstname + " " + e.Lastname
}

func (x *Student) getFullName() string {
	x.age = 54
	return x.Firstname + " " + x.Lastname
}
