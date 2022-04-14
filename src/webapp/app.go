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

func handleRequests() {

	http.HandleFunc("/", getAllCitizens)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {

	fmt.Println(fmt.Sprintf("Rest API v%d - Mux Routers listening on port %s", VERSION, port))

	handleRequests()
}
