package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strings"
	"github.com/kataras/go-errors"
	"strconv"
)

type HackExecutor struct {
	TimeApp        string
	TypeCheckApp   string
	PHPexeApp      string
	HHVMexeApp     string
	FileName       string
	CurrDirectory  string
	RedirectStdout string
	RedirectStderr string
}

type MeasuredTime struct {
	RealTime float64 `json:"real"`
	UserTime float64 `json:"user"`
	SysTime  float64 `json:"sys"`
}

// NewSomething create new instance of Something
func NewHackExecutor(filename, outfilename, currentDir string) HackExecutor {
	res := HackExecutor{
		TimeApp:        "time",
		TypeCheckApp:   "hh_client",
		PHPexeApp:      "php",
		HHVMexeApp:     "hhvm",
		FileName:       filename,
		CurrDirectory:  currentDir,
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
	// executing a hacklang program, measure the running time
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

	output, _ := ioutil.ReadAll(stdout)
	res := string(output)
	res = strings.Replace(res, he.CurrDirectory, "", -1)

	slurp, _ := ioutil.ReadAll(stderr)
	stderrOutput := string(slurp)
	et, _ := extractTime(stderrOutput)
	fmt.Println(et)
	fmt.Printf(":%s:\n", stderrOutput)

	return res, err
}

func (he HackExecutor) ExecPHP() (string, error) {
	fmt.Println(he.HHVMexeApp)
	cmd := exec.Command(he.TimeApp, he.PHPexeApp, he.FileName)

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

func extractTime(measure string) (MeasuredTime, error) {
	// extracting time information from standard error output
	// real running time
	// user program running time
	// operating system running time
	timeString := "[0-9]+\\.[0-9]+\\s+real\\s+[0-9]+\\.[0-9]+\\s+user\\s+[0-9]+\\.[0-9]+\\s+sys"
	timePattern := regexp.MustCompile(timeString)
	matched := timePattern.FindStringSubmatch(measure)

	if len(matched) != 1 {
		return MeasuredTime{}, errors.New("Cannot extract time from string")
	}

	str := string(matched[0])

	emptyPattern := regexp.MustCompile("\\s+")
	str = emptyPattern.ReplaceAllString(str, ",")

	arr := strings.Split(str, ",")
	realTime, _ := strconv.ParseFloat(arr[0], 64)
	userTime, _ := strconv.ParseFloat(arr[2], 64)
	sysTime, _ := strconv.ParseFloat(arr[4], 64)

	measuredTime := MeasuredTime{
		RealTime: realTime,
		UserTime: userTime,
		SysTime:	sysTime,
	}

	return measuredTime, nil
}
