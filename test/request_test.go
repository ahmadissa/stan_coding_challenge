package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"testing"

	"github.com/ahmadissa/stan_coding_challenge/jsonData"
)

type requestFile struct {
	expectedStatusCode int
	testName           string
	requestBody        []byte
}

// getRequets from requests directory
func getRequets() []requestFile {
	requests := []requestFile{}
	filesDir := "./requests/"
	files, err := ioutil.ReadDir(filesDir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileFound := filesDir + file.Name()
		if file.IsDir() || filepath.Ext(fileFound) != ".json" {
			continue
		}
		expectedStatusCode, err := getExpectedStatusCode(file.Name())
		if err != nil {
			log.Fatal(err.Error())
		}
		jsonBytes, err := ioutil.ReadFile(fileFound)
		if err != nil {
			fmt.Print(err)
		}
		req := requestFile{}
		req.expectedStatusCode = expectedStatusCode
		req.testName = file.Name()
		req.requestBody = jsonBytes
		requests = append(requests, req)
	}
	return requests
}

//getExpectedStatusCode expecting file name as testName.expectedStatusCode.json ex: default.200.json
func getExpectedStatusCode(file string) (int, error) {
	fileNameAry := strings.Split(file, ".")
	if len(fileNameAry) != 3 {
		return 0, errors.New("[getExpectedStatusCode]:failed to get expected status, file: " + file)
	}
	status, err := strconv.Atoi(fileNameAry[1])
	if err != nil {
		return 0, err
	}
	return status, nil
}

//postRequest and compare the response status code to the expected one, also make sure the response are as expected
func postRequest(url string, reqTest requestFile) {

	client := &http.Client{}
	data := bytes.NewBuffer(reqTest.requestBody)
	req, err := http.NewRequest(http.MethodPost, url, data)
	if err != nil {
		log.Fatal(err)
	}
	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	// make sure it returns the expected status code
	if response.StatusCode != reqTest.expectedStatusCode {
		log.Fatal("[postRequest]: unexpected status code:" + reqTest.testName)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	switch response.StatusCode {
	case http.StatusOK: // if OK is response should be in jsonData.ResponseOK stucture
		responseOK := jsonData.ResponseOK{}
		err := json.Unmarshal(responseData, &responseOK)
		if err != nil {
			fmt.Println("[postRequest]: Response is not valid:" + reqTest.testName)
			log.Fatal(err)
		}

	case http.StatusBadRequest: //// if Failed is response should be in jsonData.ResponseError stucture
		responseErr := jsonData.ResponseError{}
		err := json.Unmarshal(responseData, &responseErr)
		if err != nil {
			log.Fatal(err)
		}
		if len(responseErr.Error) == 0 {
			log.Fatal("[postRequest]: responseErr is empty testName:" + reqTest.testName)
		}
	}

}
func TestWebService(t *testing.T) {
	requests := getRequets()
	const serverURL = "http://localhost:19090" // you need to change this if you want to test different web service
	for i := range requests {
		postRequest(serverURL, requests[i])
	}

}
