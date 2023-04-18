package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/mattn/go-sqlite3"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func TestEventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println(v.Message.GetConversation())

	}
}

func main() {
	dbLog := waLog.Stdout("Database", "INFO", true)
	clientLog := waLog.Stdout("Client", "INFO", true)
	_ = sqlite3.SQLITE_REAL // dummy to import sqlite3

	// Load variables
	err := godotenv.Load("conf.env")

	if err != nil {
		clientLog.Warnf(
			"Missing conf.env trying to lookup into current environment for variables..",
		)
	}

	dbPath := os.Getenv("DB_PATH")
	adminId := os.Getenv("ADMIN")
	groupId := os.Getenv("GROUP")
	apiKey := os.Getenv("API_KEY")

	container, err := sqlstore.New("sqlite3", dbPath, dbLog)
	if err != nil {
		panic(err)
	}

	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		panic(err)
	}

	client := whatsmeow.NewClient(deviceStore, clientLog)
	lolClient := NewLolClient(apiKey)
	clientHandlers := NewClientHandler(
		client,
		lolClient,
		groupId,
		adminId,
	)
	clientHandlers.AddEventHandlers()

	if client.Store.ID == nil {
		// new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			panic(err)
		}

		for event := range qrChan {
			if event.Event == "code" {
				qrterminal.GenerateHalfBlock(event.Code, qrterminal.L, os.Stdout)
			} else {
				fmt.Println("Login event: ", event.Event)
			}
		}
	} else {
		err = client.Connect()
		if err != nil {
			panic(err)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()

}
