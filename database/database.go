package database

import (
	"log"

	"github.com/LFSCamargo/go-rest-boilerplate/env"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Connect - Is the function that connects to the mongoDB
func Connect() {
	config := env.GetConfig()
	err := mgm.SetDefaultConfig(nil, config.MongoDatabase, options.Client().ApplyURI(config.MongoHost))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to Mongo at %s", config.MongoHost)
}
