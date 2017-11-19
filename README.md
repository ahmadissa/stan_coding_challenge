stan coding challenge
===================


This simple web service done as per requirements mentioned in:
https://challengeaccepted.streamco.com.au

----------


Running the code
-------------

just run the below 
```
$ go get -u github.com/ahmadissa/stan_coding_challenge
$ cd $GOPATH/src/github.com/ahmadissa/stan_coding_challenge
$ go run main.go
```
----------

Testing
-------------

tests are available under "test" directory, the script will read rquests from json files available under "requests"

JSON files names should be formated as below:

testName.expectedStatusCode.json

example:

default.200.json



to run the test just go to test directory and run the below
```
go test
```

to add more tests just add JSON files following the above instructions

to test another web service running on remote server just change the "serverURL" in "request_test.go" file
