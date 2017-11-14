package requestHandler

import (
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
		response.Header().Set("Content-Type", "application/json")

		response.WriteHeader(status)

		log.Printf("[requestHandler]:HandlePut: sending data:" + string(responseJSON))
		response.Write(responseJSON)
	default:
		response.WriteHeader(http.StatusMethodNotAllowed)
		response.Write([]byte("Method Not Allowed"))
		return
	}

}
