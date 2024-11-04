package main

import (
	"fmt"
	"log"
	"server/internal/server"

	"github.com/labstack/echo/v4"
)

func main() {

	echo := echo.New()

	// MongoDB configuration
	config := server.Config{
		Host:     "localhost",
		Port:     27018,
		Username: "user",
		Password: "pass",
		Database: "database",
	}

	// Initialize DB connection
	mongoDB, err := server.NewDB(config)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	defer mongoDB.Close()

	// Access a collection
	collection := mongoDB.GetCollection("your_collection")
	fmt.Println("Connected to MongoDB and accessed collection:", collection.Name())

	client := &server.Server{}
	client.SetupAndLaunch(echo)

}
