package main

import (
	"flag"
	"github.com/has1985/bot_useful/clients/telegram"
	"log"
)

const (
	tgBotHost = "api.telegram.org" //make it a configurable parameter
)

func main() {

	tgClient := telegram.New(tgBotHost, mastToken())

	//fetcher := fetcher.New()

	//processor := processor.New()

	//consumer.Start(fetcher,processor)
}

func mastToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot",
	)

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
