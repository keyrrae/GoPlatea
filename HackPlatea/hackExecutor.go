package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"runtime"
	"errors"
)

type HackExecutor struct {
	TimeApp        string
	TypeCheckApp   string
	PHPexeApp      []string
	HHVMexeApp     string
	FileName       string
	CurrDirectory  string
	RedirectStdout string
	RedirectStderr string
}

type MeasuredTime struct {
	RealTime string `json:"real"`
	UserTime string `json:"user"`
	SysTime  string `json:"sys"`
}

type TaskResult struct {
	Name   string  `json:"name"`
	Output string  `json:"output"`
	Time   MeasuredTime `json:"time"`
}

var HHVM string = "hhvm"
var PHP70 string = "php7.0"

// NewSomething create new instance of Something
func NewHackExecutor(filename, outfilename, currentDir string) HackExecutor {
	res := HackExecutor{
		TimeApp:        "time",
		TypeCheckApp:   "hh_client",
		PHPexeApp:      []string{ "php", "php5.6", "php7.0" },
		HHVMexeApp:     "hhvm",
		FileName:       filename,
		CurrDirectory:  currentDir,
		RedirectStdout: "2>" + outfilename,
		RedirectStderr: "3>2",
	}

	return res
}

func (he HackExecutor) TypeCheck() (TaskResult, error) {
	// perform type checking on the program
	cmd := exec.Command(he.TypeCheckApp, he.FileName)

	output, _ := cmd.CombinedOutput()
	typeCheckoutput := string(output)
	typeCheckoutput = strings.Replace(typeCheckoutput, he.CurrDirectory, "", -1)

	res := TaskResult{
		Name: HHVM,
		Output: typeCheckoutput,
		Time: MeasuredTime{},
	}

	return res, nil
}

func (he HackExecutor) ExecHHVM() TaskResult {
	// executing a hacklang program, measure the running time
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

	output, _ := ioutil.ReadAll(stdout)
	execOutput := string(output)
	fmt.Println(execOutput)
	execOutput = strings.Replace(execOutput, he.CurrDirectory, "", -1)

	slurp, _ := ioutil.ReadAll(stderr)
	stderrOutput := string(slurp)
	exeTime, _ := extractTime(stderrOutput)

	res := TaskResult{
		Name: HHVM,
		Output: execOutput,
		Time: exeTime,
	}

	return res
}

func (he HackExecutor) ExecPHP() []TaskResult {
	var results []TaskResult
	for _, phpExeApp := range he.PHPexeApp {
		whichCommand := exec.Command("which", phpExeApp)

		whichRes, _ := whichCommand.CombinedOutput()
		if string(whichRes) != "" {
			cmd := exec.Command(he.TimeApp, phpExeApp, he.FileName)

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

			output, _ := ioutil.ReadAll(stdout)
			execOutput := string(output)
			execOutput = strings.Replace(execOutput, he.CurrDirectory, "", -1)

			slurp, _ := ioutil.ReadAll(stderr)
			stderrOutput := string(slurp)

			exeTime, _ := extractTime(stderrOutput)
			res := TaskResult{
				Name: PHP70,
				Output: execOutput,
				Time: exeTime,
			}
			results = append(results, res)
		}
	}

	return results
}

func extractTime(measure string) (MeasuredTime, error) {
	// extracting time information from standard error output
	// real execution time
	// user program execution time
	// operating system execution time
	var timeString string
	switch runtime.GOOS {
	case "darwin":
		timeString = "[0-9]+\\.[0-9]+\\s+real\\s+[0-9]+\\.[0-9]+\\s+user\\s+[0-9]+\\.[0-9]+\\s+sys"
	case "linux":
		timeString = "real\\s+[0-9]+\\.[0-9]+\\s+user\\s+[0-9]+\\.[0-9]+\\s+sys\\s+[0-9]+\\.[0-9]+\\s+"
	}

	timePattern := regexp.MustCompile(timeString)
	matched := timePattern.FindStringSubmatch(measure)

	if len(matched) != 1 {
		return MeasuredTime{}, errors.New("Cannot extract time from string")
	}

	str := string(matched[0])

	emptyPattern := regexp.MustCompile("\\s+")
	str = emptyPattern.ReplaceAllString(str, ",")

	arr := strings.Split(str, ",")
	realTime := arr[0]
	userTime := arr[2]
	sysTime := arr[4]

	measuredTime := MeasuredTime{
		RealTime: realTime,
		UserTime: userTime,
		SysTime:	sysTime,
	}

	return measuredTime, nil
}
