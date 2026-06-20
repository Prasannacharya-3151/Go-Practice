package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//person struct
type Person struct{
	Name string `json:"name"`
	Age int32   `json:"age"`
}

//sample data
var personData = map[string]Person{
	"1":{Name:"john doe", Age:30},
	"2":{Name:"john doe", Age:28},
	"3":{Name:"john doe", Age:25},
}

//handler function for the endpoint
func getPersonHandler(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
	}

	//check is the id exist in the personDat map
	person, exists := personData[id]
	if !exists {
		http.Error(w, "Person not found", http.StatusNotFound)
		return
	}

	//set the response header to appliction/json
	w.Header().Set("Content-Type", "application/json")

	//enocode the person data to json and write to the response
	if err := json.NewEncoder(w).Encode(person); err !=nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func main(){
	//define the port
	port := 8080

	//print the confirmation message
	fmt.Printf("Server started on the port %d\n", port)

	//set up the endpoint and the handler function
	http.HandleFunc("/person", getPersonHandler)

	//start the server
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}