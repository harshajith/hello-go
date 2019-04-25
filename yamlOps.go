package main

import (
	"fmt"
	"log"

	"gopkg.in/yaml.v2"
)

var data = `
a: Easy!
b:
  c: 2
  d: [3, 4]
`

var vpn = `
---
name: vpn-demo
owners: ["457785", "4578585"]
ports:
  t1: 4333
  d1: 4233
  p1: 4334
`

type V1 struct {
	Name   string
	Owners []string
	Ports  map[interface{}]interface{}
}

type T struct {
	A string
	B struct {
		RenamedC int   `yaml:"c"`
		D        []int `yaml:",flow"`
	}
}

func yamlOps() {

	t := T{}

	err := yaml.Unmarshal([]byte(data), &t)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	fmt.Printf("--- t:\n%v", t)

	v := V1{}

	err1 := yaml.Unmarshal([]byte(vpn), &v)
	if err1 != nil {
		log.Fatalf("error: %v", err1)
	}
	fmt.Printf("--- t:\n%v\n\n", v)

}
