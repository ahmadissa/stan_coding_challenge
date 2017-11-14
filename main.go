package main

import (
	"log"
	"net/http"

	"github.com/ahmadissa/stan_coding_challenge/requestHandler"
)

func main() {
	http.HandleFunc("/", requestHandler.HandlePut)
	// open port and wait for requests
	err := http.ListenAndServe(":19090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
