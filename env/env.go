package env

import "os"

const defaultPort = ":3000"
const defaultMongoURI = "mongodb://localhost:27017"
const defaultMongoDatabase = "boilerplate"
const defaultJwt = "top secret"

// Config - All the configuration for the application to work
type Config struct {
	Port          string
	MongoHost     string
	MongoDatabase string
	Jwt           string
}

// GetConfig - Gets all the env variables
func GetConfig() Config {
	mongoURI := os.Getenv("MONGO_URI")
	database := os.Getenv("MONGO_DATABASE")
	port := os.Getenv("PORT")
	jwt := os.Getenv("JWT")

	if mongoURI == "" && database == "" && port == "" && jwt == "" {
		port = defaultPort
		mongoURI = defaultMongoURI
		database = defaultMongoDatabase
		jwt = defaultJwt
	}

	return Config{
		Port:          port,
		MongoHost:     mongoURI,
		MongoDatabase: database,
		Jwt:           jwt,
	}
}
