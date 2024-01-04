package main

import (
	"flag"
	"log"
)

func main() {
	t:= mastToken() //token

	//tgClient := telegram.New(token)

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
