package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/strata-io/service-extension/orchestrator"
)


type User []struct { //structure for user info 
	ID       	int    `json:"id"`
	Email    	string `json:"email"`
}

func CreateEmailHeader(api orchestrator.Orchestrator, _ http.ResponseWriter, _ *http.Request) (http.Header, error) {
	logger := api.Logger()
	logger.Debug("se", "building email custom header")
	
	session, err := api.Session()
	if err != nil {
		logger.Error("se", "unable to retrieve session", "error", err.Error())
		return nil, err
	}

    url := "https://jsonplaceholder.typicode.com/users?id=2" //define url to retrieve user information by getting userId input
	response, err := http.Get(url) //checking url is up and retrieve user information, if it fails throw an error
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() //close connection once info is read

 	body, err := ioutil.ReadAll(response.Body) //read GET response from server
	if err != nil { //check for error
		log.Fatal(err)
	}
    
	var userInfo User //declare userInfo variable and assign User struct to it
	json.Unmarshal(body, &userInfo)

	fmt.Println(userInfo[0].Email) //display GET response for email, in an array 0 since it returns only one entry

    email := userInfo[0].Email
    header := make(http.Header)
    header["CUSTOM-EMAIL"] = []string{email}
    return header, nil
}

func main() {
	CreateEmailHeader()
}