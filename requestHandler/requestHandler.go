package requestHandler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmadissa/stan_coding_challenge/dataProcessor"
)

//HandlePut will handle requests and validate the request and respond accordingly
func HandlePut(response http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	method := req.Method
	switch method {
	case "POST":
		responseJSON, status := dataProcessor.Process(req.Body)
		response.WriteHeader(status)

		log.Printf("[requestHandler]:HandlePut: sending data:" + responseJSON)
		fmt.Fprintf(response, responseJSON)
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(response, "Method Not Allowed")
		return
	}

}
