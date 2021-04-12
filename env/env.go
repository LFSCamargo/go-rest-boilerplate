package env

import "os"

const defaultPort = ":3000"
const defaultMongoURI = "mongodb://localhost:27017"
const defaultMongoDatabase = "trivia-api"
const defaultJwt = "top secret"
const defaultEnv = "DEV"

// Config - All the configuration for the application to work
type Config struct {
	Port          string
	MongoHost     string
	MongoDatabase string
	Jwt           string
	Env           string
}

// GetConfig - Gets all the env variables
func GetConfig() Config {
	mongoURI := os.Getenv("MONGO_URI")
	database := os.Getenv("MONGO_DATABASE")
	port := os.Getenv("PORT")
	jwt := os.Getenv("JWT")
	environment := os.Getenv("ENV")

	if mongoURI == "" && database == "" && port == "" && jwt == "" {
		port = defaultPort
		mongoURI = defaultMongoURI
		database = defaultMongoDatabase
		jwt = defaultJwt
		environment = defaultEnv
	}

	return Config{
		Port:          port,
		MongoHost:     mongoURI,
		MongoDatabase: database,
		Jwt:           jwt,
		Env:           environment,
	}
}
