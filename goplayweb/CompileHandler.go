package main

import (
	"net/http"
)

func init() {
	http.HandleFunc("/compile", compileHandler)
}

func compileHandler(w http.ResponseWriter, r *http.Request) {

}