package main

import "net/http"

func init() {
	http.HandleFunc("/fmt", fmtHandler)
}

func fmtHandler(w http.ResponseWriter, r *http.Request){

}