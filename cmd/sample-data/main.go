package main

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws/endpoints"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/rejlersembriq/hooked/pkg/dynamo"
	"github.com/rejlersembriq/hooked/pkg/participant"
	"log"
)

const (
	profile   = "larwef"
	region    = endpoints.EuWest1RegionID
	tableName = "hooked-participants"
)

func main() {
	conf, err := external.LoadDefaultAWSConfig(external.WithSharedConfigProfile(profile))
	if err != nil {
		log.Fatalf("Message getting AWS config: %v", err)
	}
	conf.Region = region

	ddb := dynamodb.New(conf)

	repo := dynamo.New(ddb, tableName)

	p := &participant.Participant{
		//Id:      "5991dd0d-da1e-4088-9889-72067ab9d467",
		Email:   "larwef@gmail.com",
		Name:    "Lars Wefald",
		Phone:   "12345678",
		Org:     "Rejlers Embriq",
		Comment: "test",
		Score:   1,
	}

	res, err := repo.Save(p)
	if err != nil {
		log.Fatalf("Error saving participant: %v", err)
	}

	bytes, err := json.MarshalIndent(res, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling result: %v", err)
	}

	fmt.Printf("Save:\n%s\n", string(bytes))

	get, err := repo.Get(res.Id)
	if err != nil {
		log.Fatalf("Error getting participant: %v", err)
	}

	bytes, err = json.MarshalIndent(get, "", "    ")
	if err != nil {
		log.Fatalf("Error marshalling result: %v", err)
	}

	fmt.Printf("Get:\n%s\n", string(bytes))

	participants, err := repo.GetAll()
	if err != nil {
		log.Fatalf("Error getting participans: %v", err)
	}
	for i, p := range participants {
		bytes, err = json.MarshalIndent(p, "", "    ")
		if err != nil {
			log.Fatalf("Error marshalling result: %v", err)
		}

		fmt.Printf("Get all %d:\n%s\n", i, string(bytes))
	}
}
