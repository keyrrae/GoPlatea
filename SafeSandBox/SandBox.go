package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
	"path/filepath"
	"go/token"
)

const maxRunTime = 5 * time.Second

func main() {
	http.HandleFunc("/compile", compileHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func compileHandler(w http.ResponseWriter, req *http.Request) {
	// Compile and run request handler
	var compileReq CompileRunRequest
	if req.PostFormValue("version") == "2" {
		compileReq.Body = req.PostFormValue("body")
	} else {
		err := json.NewDecoder(req.Body).Decode(&req)
		if err != nil {
			http.Error(w, fmt.Sprintf("error decoding request: %v", err), http.StatusBadRequest)
			return
		}
	}
	// TODO: call compile and run and handle errors
}

func compileAndRun(req *CompileRunRequest) (*CompileRunResponse, error) {
	// First, create a temporary folder to put the sandbox
	tempDirectory, err := ioutil.TempDir("", "sandbox")
	if err != nil {
		return nil, fmt.Errorf("error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tempDirectory)

	in := filepath.Join(tempDirectory, "main.go")
	err = ioutil.WriteFile(in, []byte(req.Body), 0440)
	if err != nil {
		return nil, fmt.Errorf("Error when creating a temporary file %v: %v", in, err)
	}

	fileSet := token.NewFileSet()
	// TODO: body of compile and run
	_ = fileSet

	return &CompileRunResponse{Head: "head", Body: "body"}, nil
}

func checkTimeout(command *exec.Cmd, duration time.Duration) error {
	if err := command.Start(); err != nil {
		return err
	}
	errChannel := make(chan error, 1)

	go func() {
		errChannel <- command.Wait()
	}()

	timer := time.NewTimer(duration)
	select {
	case err := <-errChannel:
		timer.Stop()
		return err
	case <-timer.C:
		command.Process.Kill()
		timeoutErr := errors.New("compile and run timed out")
		return timeoutErr
	}
}
