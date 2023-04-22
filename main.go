package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mau.fi/whatsmeow/types/events"
)

func TestEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println(v.Message.GetConversation())

	}
}

func main() {
	// Load variables
	err := godotenv.Load("conf.env")

	if err != nil {
		fmt.Println(
			"Missing conf.env trying to lookup into current environment for variables..",
		)
	}

	dbPath := os.Getenv("DB_PATH")
	adminId := os.Getenv("ADMIN")
	groupId := os.Getenv("GROUP")
	apiKey := os.Getenv("API_KEY")

	wppClient := NewWppClient(dbPath, apiKey)
	lolClient := NewLolClient(apiKey)
	leviBot := NewLeviClient(wppClient, lolClient, groupId, adminId)

	wppClient.AddEventHandler(leviBot.CommandHandler)
	leviBot.CheckForNewMatches()
}
