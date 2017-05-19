package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	//"encoding/json"
	"encoding/json"
)

var hackExectr HackExecutor
var envHelper EnvHelper

type respBody struct {
	Result string `json:"result"`
}

type taskResult struct {
	Name   string  `json:"name"`
	Output string  `json:"output"`
	Status string  `json:"status"`
	Time   float64 `json:"time"`
}

type ExeReq struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func init() {
	envHelper = NewEnvHelper([]string{"test.php", "output"})
	hackExectr = NewHackExecutor("test.php", "output", envHelper.GetCurrDirectory())
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HacklangHandler).Methods("POST")

	port := "" //os.Getenv("PORT")
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
		http.Error(w, err.Error()+str, http.StatusInternalServerError)
	}

	body, err := ioutil.ReadAll(r.Body)

	fmt.Println(string(body))
	exeReq := ExeReq{}
	json.Unmarshal(body, &exeReq)

	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "can't read body", http.StatusBadRequest)
		return
	}

	//err = envHelper.WriteProgToSystem(string(body))
	err = envHelper.WriteProgToSystem(exeReq.Code)
	if err != nil {
		log.Printf("Error creating program: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	str, err = hackExectr.TypeCheck()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Type check error:\n"+str)
		return
	}

	str, err = hackExectr.ExecHHVM()
	if err != nil {
		w.Header().Set("Content-Type", "text/plain")

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "%s", "Execution error:\n"+str)
		return
	}
	/*
		str, err = hackExectr.ExecPHP()
		if err != nil {
			w.Header().Set("Content-Type", "text/plain")

			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s", "Execution error:\n" + str)
			return
		}
	*/
	w.Header().Set("Content-Type", "text/plain")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", str)
}
