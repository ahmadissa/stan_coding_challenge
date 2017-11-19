package main

import (
	"log"
	"net/http"

	"github.com/ahmadissa/stan_coding_challenge/requestHandler"
)

const host = ":19090"

func main() {
	http.HandleFunc("/", requestHandler.HandlePost)
	// open port and wait for requests
	err := http.ListenAndServe(host, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
