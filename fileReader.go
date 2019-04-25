package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

// type vpnRegistry struct {
// 	entries []vpnConfig
// }

type V struct {
	Name   string
	Owners []string
	Ports  map[interface{}]interface{}
}

func readFile() {
	resp, err := http.Get("https://github.com/harshajith/hello-go/blob/master/conf/fss-registry.yml?raw=True")
	if err != nil {

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	vpnReg := V{}

	err = yaml.Unmarshal(body, &vpnReg)
	if err != nil {
		log.Print(err)
	}

	fmt.Println(vpnReg)

}
