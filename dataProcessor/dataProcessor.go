package dataProcessor

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/ahmadissa/stan_coding_challenge/jsonData"
)

func validateRequestShow(requestShow jsonData.RequestShow) error {
	if len(requestShow.Slug) == 0 {
		return errors.New(" Slug is empty")
	}
	if len(requestShow.Title) == 0 {
		return errors.New(" Title is empty")
	}
	if len(requestShow.Image.ShowImage) == 0 {
		return errors.New(" ShowImage is empty")
	}
	return nil
}

//getResponse return error if request is not valid, or return response
func getResponse(request jsonData.Request, response *jsonData.ResponseOK) error {
	if len(request.Payload) == 0 { // check if payload is empty
		return errors.New(" payload is empty")
	}
	for i := range request.Payload {
		if !request.Payload[i].Drm || request.Payload[i].EpisodeCount == 0 { // if not needed dont validate
			continue
		}
		if err := validateRequestShow(request.Payload[i]); err != nil {
			return err
		}
		show := jsonData.ResponseShow{}
		show.Image = request.Payload[i].Image.ShowImage
		show.Slug = request.Payload[i].Slug
		show.Title = request.Payload[i].Title
		response.Response = append(response.Response, show)

	}
	return nil
}
func getError(errString string) string {
	errResponse := jsonData.ResponseError{}
	errResponse.Error = "Could not decode request: " + errString
	responseJSON, _ := json.Marshal(errResponse)
	return string(responseJSON)
}

//Process request body, validate request and reply with response string and status code
func Process(body io.ReadCloser) (string, int) {
	jsonBytes, err := ioutil.ReadAll(body)

	if err != nil { // if failed to decode return error
		return getError("Reading Request failed"), http.StatusBadRequest
	}
	request := jsonData.Request{}
	err = json.Unmarshal(jsonBytes, &request)
	if err != nil { // if failed to decode return error
		return getError("JSON parsing failed"), http.StatusBadRequest
	}
	defer body.Close()
	// check if request if valid
	response := jsonData.ResponseOK{}
	response.Response = []jsonData.ResponseShow{}
	err = getResponse(request, &response)
	if err != nil { // if not valid
		return getError(err.Error()), http.StatusBadRequest
	}
	resposeBytes, _ := json.Marshal(response)
	return string(resposeBytes), http.StatusOK
}
