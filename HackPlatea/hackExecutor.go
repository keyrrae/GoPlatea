package main

import (
	"os/exec"
	"strings"
	"fmt"
	"log"
	"io/ioutil"
)

type HackExecutor struct {
	TimeApp	string
	TypeCheckApp  string
	PHPexeApp	string
	HHVMexeApp  string
	FileName      string
	CurrDirectory string
	RedirectStdout string
	RedirectStderr string
}

// NewSomething create new instance of Something
func NewHackExecutor(filename, outfilename, currentDir string) HackExecutor {
	res := HackExecutor{
		TimeApp:	"time",
		TypeCheckApp:  "hh_client",
		PHPexeApp:	"php",
		HHVMexeApp:  "hhvm",
		FileName:      filename,
		CurrDirectory: currentDir,
		RedirectStdout: "2>" + outfilename,
		RedirectStderr: "3>2",
	}

	return res
}

func (he HackExecutor) TypeCheck() (string, error) {
	// perform type checking on the program
	cmd := exec.Command(he.TypeCheckApp, he.FileName)

	output, err := cmd.CombinedOutput()
	res := string(output)
	res = strings.Replace(res, he.CurrDirectory, "", -1)
	if err != nil {
		return res, err
	}

	return res, nil
}

func (he HackExecutor) ExecHHVM() (string, error) {
	fmt.Println(he.HHVMexeApp)
	cmd := exec.Command(he.TimeApp, he.HHVMexeApp, he.FileName)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	fmt.Printf(":%s:\n", slurp)

	output, _ := ioutil.ReadAll(stdout)
	res := string(output)

	res = strings.Replace(res, he.CurrDirectory, "", -1)
	if err != nil {
		return res, err
	}

	return res, nil
}
