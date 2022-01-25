package graph

import (
	"log"
	"os"

	"github.com/Peyton232/freelance-auction/database"
	"github.com/joho/godotenv"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

var db = database.Connect(goDotEnvVariable("MONGODB_URI"), goDotEnvVariable("MONGODB_NAME"))

// use godot package to load/read the .env file and
// return the value of the key
func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
