package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
	"os"
)

var hackExectr HackExecutor
var envHelper EnvHelper
var nameMap map[string]string

type respBody struct {
	Result string `json:"result"`
}

type ExeReq struct {
	Language string `json:"language"`
	Code     string `json:"code"`
}

func init() {
	envHelper = NewEnvHelper([]string{"test.php", "output"})
	hackExectr = NewHackExecutor("test.php", "output", envHelper.GetCurrDirectory())
	nameMap = map[string]string {
		"hhvm": "HHVM",
		"php":   "PHP",
		"php5.6": "PHP 5.6",
		"php7.0": "PHP 7.0",
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", HacklangHandler).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		// default port number is 8000
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

	err = envHelper.WriteProgToSystem(exeReq.Code)
	if err != nil {
		log.Printf("Error creating program: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res []TaskResult

	if exeReq.Language == "php" {
		phpExeRes := hackExectr.ExecPHP()
		res = append(res, phpExeRes...)
	}

	typeCheckRes, err := hackExectr.TypeCheck()
	if err != nil {
		// type check error
		res = append(res, typeCheckRes)
	} else {
		// no type check error, execute the program
		hhvmExeRes := hackExectr.ExecHHVM()
		res = append(res, hhvmExeRes)
	}

	resJ, _ := json.Marshal(res)
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", resJ)
}
