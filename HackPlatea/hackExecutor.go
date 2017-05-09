package main

import (
	"os/exec"
	"strings"
)

type HackExecutor struct {
	TypeCheckApp  string
	ExecutionApp  string
	FileName      string
	CurrDirectory string
}

// NewSomething create new instance of Something
func NewHackExecutor(filename, currentDir string) HackExecutor {
	res := HackExecutor{
		TypeCheckApp:  "hh_client",
		ExecutionApp:  "hhvm",
		FileName:      filename,
		CurrDirectory: currentDir,
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

func (he HackExecutor) ExecProgram() (string, error) {
	cmd := exec.Command(he.ExecutionApp, he.FileName)

	output, err := cmd.CombinedOutput()
	res := string(output)
	res = strings.Replace(res, he.CurrDirectory, "", -1)
	if err != nil {
		return res, err
	}

	return res, nil
}
