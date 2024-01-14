package main

import (
	"flag"
	tgClient "github.com/has1985/bot_useful/clients/telegram"
	event_consumer "github.com/has1985/bot_useful/consumer/event-consumer"
	"github.com/has1985/bot_useful/events/telegram"
	"github.com/has1985/bot_useful/storage/files"
	"log"
)

const (
	tgBotHost   = "api.telegram.org" //make it a configurable parameter
	storagePath = "file_storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mastToken()),
		files.New(storagePath),
	)

	log.Printf("service started")
	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatalf("service is stopped", err)
	}
}

func mastToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
