package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

const (
	VERSION      int    = 4
	port         string = "8000"
	DYNAMO_TABLE string = "ur-u-table"
	AWS_REGION   string = "eu-south-1"
)

type Citizen struct {
	id         string
	CF         string
	FirstName  string
	LastName   string
	ApiVersion int
}

func getAllCitizens(w http.ResponseWriter, r *http.Request) {

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Region: aws.String(AWS_REGION)})
	table := db.Table(DYNAMO_TABLE)

	// get all items
	var results []Citizen
	err := table.Scan().All(&results)

	if err != nil {
		fmt.Println("Error gettings items ", err)
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
