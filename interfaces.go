package main

type Department interface {
	name() string
	dean() string
}

type Lab interface {
	totalFacilities() int
}

type CSE struct {
	description string
}

func (d CSE) name() string {
	return "CSE"
}

func (d CSE) dean() string {
	return "Gunathilaka"
}

func (d CSE) totalFacilities() int {
	return 10
}
