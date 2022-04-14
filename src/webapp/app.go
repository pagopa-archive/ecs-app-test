package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	VERSION int    = 4
	port    string = "8060"
)

type Citizen struct {
	id         string
	CF         string
	FirstName  string
	LastName   string
	ApiVersion int
}

func getAllCitizens(w http.ResponseWriter, r *http.Request) {

	// get all items
	var results []Citizen

	results = []Citizen{
		{
			id:         "001",
			CF:         "AAAAAA",
			FirstName:  "Tony",
			LastName:   "Manero",
			ApiVersion: VERSION,
		},
	}

	// Print all http headers:
	for key, val := range r.Header {
		fmt.Println(fmt.Sprintf("Header %s  Value %s", key, val))
	}

	json.NewEncoder(w).Encode(results)
}

func createCitizen(w http.ResponseWriter, r *http.Request) {
	// Declare a new Person struct.
	var p Citizen

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the Person struct...
	fmt.Fprintf(w, "Person: %+v", p)
	fmt.Println("Person: ", p)
}

func handleRequests() {

	mux := http.NewServeMux()
	mux.HandleFunc("/person/create", createCitizen)
	mux.HandleFunc("/person/read", getAllCitizens)

	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}

func main() {

	fmt.Println(fmt.Sprintf("Rest API v%d - Mux Routers listening on port %s", VERSION, port))

	handleRequests()
}
