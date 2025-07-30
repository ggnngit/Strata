package main

import (
	"encoding/json" //function to encode/decode JSON
	"fmt"           //allow for fmt write out
	"io/ioutil"     //add io/util functions
	"log"           //add logging functions
	"net/http"      //server functions
)

type User []struct { //structure for user info 
	ID       	int    `json:"id"`
	Email    	string `json:"email"`
	Phone 		string `json:"phone"`
	Company struct {
		Name        string `json:"name"`
		CatchPhrase string `json:"catchPhrase"`
		Bs          string `json:"bs"`
	} `json:"company"`
}

func PrettyPrint(i interface{}) string { //pretty print to show name and value in JSON; else it only shows value after unmarshall
    pPrint, _ := json.MarshalIndent(i, "", "\t")
    return string(pPrint)
}

func GetUser (){
	url := "https://jsonplaceholder.typicode.com/users" //define url to retrieve user information
	response, err := http.Get(url) //checking url is up and retrieve user information, if it fails throw an error
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() //close connection once info is read

 	body, err := ioutil.ReadAll(response.Body) //read GET response from server
	if err != nil { //check for error
		log.Fatal(err)
	}
	// fmt.Println(string(body)) //just a sanity check that it's working and also prints out user details

	var userInfo User //declare userInfo variable and assign User struct to it

	if jsonErr := json.Unmarshal(body, &userInfo); jsonErr != nil { //check for parsing error
		fmt.Printf("Couldn't parse JSON", jsonErr)
		return
	}
	fmt.Println(PrettyPrint(userInfo)) //display GET response for id, email, phone, and company info
}

func main() {
	GetUser() //calling GetUser function
}