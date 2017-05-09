package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
)

var hackExectr HackExecutor
var envHelper EnvHelper

func init() {
	envHelper = NewEnvHelper("test.hh")
	hackExectr = NewHackExecutor("test.hh", envHelper.GetCurrDirectory())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HacklangHandler).Methods("POST")

	port := ""//os.Getenv("PORT")
	if port == "" {
		// default port number is 5000
		port = "8000"
	}

	// Fire up the server
	err := http.ListenAndServe(":"+port, r)
	if err != nil {
		panic(err)
	}
}


func HacklangHandler(w http.ResponseWriter, r *http.Request) {
	str, err := envHelper.ClearWorkspace()
	if err != nil {
		http.Error(w, err.Error() + str, http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	err = envHelper.WriteProgToSystem(string(body))
	if err != nil {
		log.Printf("Error create program: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str, err = hackExectr.TypeCheck()
	if err != nil {
		http.Error(w, str, http.StatusNotAcceptable)
		return
	}

	str, err = hackExectr.ExecProgram()
	if err != nil {
		http.Error(w, str, http.StatusNotAcceptable)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", str)
}