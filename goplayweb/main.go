package main

import (
	"log"
	"net/http"
	"flag"
)

var	serveAddress *string

func init() {
	serveAddress = flag.String("l", ":3000", "Listen address.")
}


func main() {

	log.Printf("Serving GoPlatea at localhost:%v...\n", "3000")

	if err := http.ListenAndServe(*serveAddress, nil); err != nil {
		log.Fatal(err)
	}
}
