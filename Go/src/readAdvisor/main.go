package main

import (
	"flag"
	"log"

	tgClient "readAdvisor/clients/telegram"
	"readAdvisor/consumer/event_consumer"
	"readAdvisor/events/telegram"
	"readAdvisor/storage/files"
)

// chmod -R 775 ./Go/src/readAdvisor/files_storage/

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		files.New(storagePath),
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, 100)

	if err := consumer.Start(); err != nil {
		log.Fatal("service was stopped", err)
	}

}

func mustToken() string {
	token := flag.String(
		"tg-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("Token is not specified")
	}

	return *token
}
