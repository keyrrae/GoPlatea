package main

import (
	"os/exec"
	"os"
	"log"
	"strings"
	"fmt"
)

func getCurrDirectory() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return dir + "/"
}

func clearWorkspace(){

}

func writeProgToSystem(prog string){

}

func typeCheck() (string, error) {
	app := "hh_client"
	filename := "test.hh"

	cmd := exec.Command(app, filename)

	output, err := cmd.CombinedOutput()
	res := string(output)
	res = strings.Replace(res, getCurrDirectory(), "", -1)
	//fmt.Println(res)
	if err != nil {
		return res, err
	}

	return res, nil
}

func execProgram(){

}

func main() {
	getCurrDirectory()
	typeCheckResult, err := typeCheck()
	if err != nil {
		fmt.Print(typeCheckResult)
		return
	}
	fmt.Print(typeCheckResult)
}
