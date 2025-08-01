package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	client "github.com/kjhickman/bnet-client"
)

func main() {
	// Load environment variables from .env file in examples directory
	err := godotenv.Load("../.env")
	if err != nil {
		log.Printf("Warning: Could not load .env file: %v", err)
	}

	clientID := os.Getenv("BNET_CLIENT_ID")
	clientSecret := os.Getenv("BNET_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Please set BNET_CLIENT_ID and BNET_CLIENT_SECRET environment variables")
	}

	config := client.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	client, err := client.NewClient(config)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	token, err := client.GetWoWToken(context.Background())
	if err != nil {
		log.Fatalf("Failed to get WoW token info: %v", err)
	}

	fmt.Printf("WoW Token Price: %d gold\n", token.Price)
}
