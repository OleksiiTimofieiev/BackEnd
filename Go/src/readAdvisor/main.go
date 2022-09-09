package main

import (
	"context"
	"flag"
	"log"

	tgClient "readAdvisor/clients/telegram"
	"readAdvisor/consumer/event_consumer"
	"readAdvisor/events/telegram"
	"readAdvisor/storage/sqlite"
)

// chmod -R 775 ./Go/src/readAdvisor/files_storage/

const (
	tgBotHost         = "api.telegram.org"
	sqliteStoragePath = "/home/olekdsii/Desktop/data/sqlite/storage.db"
	batchSize         = 100
)

func main() {
	// s:= files.New(storagePath)
	s, err := sqlite.New(sqliteStoragePath)
	if err != nil {
		log.Fatalf("can`t connect to storage: ", err)
	}
	if err = s.Init(context.TODO()); err != nil {
		log.Fatalf("can`t init storage: ", err)

	}
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, mustToken()),
		s,
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
