package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	VERSION int    = 3
	port    string = "8000"
)

type Citizen struct {
	Id         string
	CF         string
	FirstName  string
	LastName   string
	ApiVersion int
}

var Citizens = []Citizen{
	{Id: "001", FirstName: "Jhon", LastName: "Smith", CF: "MRTMTT91D08F205J", ApiVersion: VERSION},
	{Id: "002", FirstName: "Robert", LastName: "De Niro", CF: "MLLSNT82P65Z404U", ApiVersion: VERSION},
}

func getAllCitizens(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Citizens)
}

func handleRequests() {

	http.HandleFunc("/", getAllCitizens)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func main() {
	fmt.Println(fmt.Sprintf("Rest API v%d - Mux Routers listening on port %s", VERSION, port))

	handleRequests()
}
