package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"
)

type V struct {
	Name   string
	Owners []string
	Ports  map[string]int
}

type ForkRepoRequest struct {
	forkFromProjectName string
	forkFromRepoName    string
	desiredProjectName  string
	desiredRepoName     string
	defaultBranch       string
	branches            []Branch
}

type Branch struct {
	branchName string
	reviewers  []string
}

func readFile() {
	resp, err := http.Get("https://github.com/harshajith/hello-go/blob/master/conf/fss-registry.yml?raw=True")
	if err != nil {

	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	vpnReg := []V{}

	err = yaml.Unmarshal(body, &vpnReg)
	if err != nil {
		log.Print(err)
	}

	for _, vpnConfig := range vpnReg {
		vpnName := vpnConfig.Name
		owners := vpnConfig.Owners
		ports := vpnConfig.Ports

		vpnRepoRequest := populateVpnRepoRequest(vpnName, owners, ports)

		fmt.Println(vpnRepoRequest)

		for key, value := range vpnConfig.Ports {
			fmt.Printf("branch name %v, port number: %v", key, value)
		}

	}
	fmt.Println(vpnReg)

}

func populateVpnRepoRequest(vpnName string, owners []string, ports map[string]int) ForkRepoRequest {
	forkRequest := ForkRepoRequest{}
	forkRequest.forkFromProjectName = "FSP"
	forkRequest.forkFromRepoName = "solace-vpn-onboarding-template"
	forkRequest.desiredProjectName = "FSS"
	forkRequest.desiredRepoName = vpnName
	forkRequest.defaultBranch = "master"

	branch := Branch{}
	branch.branchName = "master"
	branch.reviewers = []string{"1571095"}

	forkRequest.branches = []Branch{branch}
	return forkRequest
}
