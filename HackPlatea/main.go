package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"io/ioutil"
	"log"
	"fmt"
	//"encoding/json"
)

var hackExectr HackExecutor
var envHelper EnvHelper

type respBody struct {
	Result	string `json:"result"`
}

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
	w.Header().Set("Access-Control-Allow-Origin", "*")
	str, err := envHelper.ClearWorkspace()
	if err != nil {
		http.Error(w, err.Error() + str, http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))
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
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Type check error:\n" + str)
		return
	}

	str, err = hackExectr.ExecProgram()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Execution error:\n" + str)
		return
	}

	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", str)
}