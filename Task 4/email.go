package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/strata-io/service-extension/orchestrator"
	"encoding/json"
)

type User []struct { //structure for user info 
	ID       	int    `json:"id"`
	Email    	string `json:"email"`
}

func GetUser (userId int){
    url := fmt.Sprintf("https://jsonplaceholder.typicode.com/users?id=%d", userId) //define url to retrieve user information by getting userId input
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

	if jsonErr := json.Unmarshal(body, &userInfo); jsonErr != nil { //check for parsing error
		fmt.Printf("Couldn't parse JSON", jsonErr)
		return
	}
	fmt.Println(userInfo[0].Email) //display GET response for email, works in an array 0 since it returns only one entry
}

func CreateEmailHeader (
    api orchestrator.Orchestrator,
    _ http.ResponseWriter,
    _ *http.Request,
)(http.Header, error){
    logger := api.Logger()
    logger.Info("se", "building email custom header")
    session, err := api.Session()
    if err != nil {
        return nil, fmt.Errorf("unable to retrieve Orchestrator session: %w", err)
    }
}
    logger.Debug("se", "retrieving email from mock endpoint..")

    email := userInfo[0].email
    header := make(http.Header)
    header["CUSTOM-EMAIL"] = []string{email}

    return header, nil

func main (){
    fmt.Println("Input user id")
    var userId int
    _, err := fmt.Scanf("%d", &userId)
    if err != nil {
        log.Fatal(err)
    }
    GetUser(userId)
}