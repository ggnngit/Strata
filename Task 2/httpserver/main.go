package main

import (
	"encoding/json" //function to encode/decode JSON
	"fmt"           //allow for fmt write out
	"net/http"      //server functions
)

type Ping struct { //create json struct for /ping
	Message string `json:"message"`
}

type Echo struct { //create json struct for /echo
	Name string `json:"name"`
}

func ping(w http.ResponseWriter, r *http.Request) { //create function to receive GET requests
	ping := Ping{Message: "pong"} //initialize ping variable and assign Ping struct
	pJson, err := json.Marshal(ping) //marshalling to convert data into JSON from pJson variable
	if r.Method != http.MethodGet{ //allow only GET requests
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("HTTP method %q not allowed", r.Method))) //display error message
		return
	}
	if err != nil { //handling error in case of a bad request 
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write(pJson) //json response to GET request
}

func echo(w http.ResponseWriter, r *http.Request) { //create function to receive POST requests
	if r.Method != http.MethodPost{ //allow only POST requests
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte(fmt.Sprintf("HTTP method %q not allowed", r.Method))) //display error message
		return
	}
	var e Echo //declare e variable and assign to Echo struct
	err := json.NewDecoder(r.Body).Decode(&e) //retrieve POST information from body 
	if err != nil{  //bad request handling
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	eJson, err := json.Marshal(e) //marshalling to convert data into JSON from eJson variable
	if err != nil{
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json") //convert header to JSON
	w.Write(eJson) //write out response in JSON format
}

func main() {
	http.HandleFunc("/ping", ping) //call ping function to serve /ping endpoint
	http.HandleFunc("/echo", echo) //call echo function to serve /echo endpoint
	http.ListenAndServe(":8080", nil) //initialize http server on port 8080
}
