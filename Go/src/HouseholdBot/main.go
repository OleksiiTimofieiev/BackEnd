package main

import (
	// "fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botToken = "5547700494:AAHBQS1NUlsSpHS4Ob6VbQkWRjwz8ov7smw"
)

func processCommand(bot *tgbotapi.BotAPI, command string, update tgbotapi.Update) {
	if len(command) >= 0 {
		// fmt.Println(command)

		if command == "status" {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Електрика в домі працює")

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken)

	if err != nil {
		log.Panic(err)
	}

	bot.Send(tgbotapi.NewSetMyCommands(
		// tgbotapi.BotCommand{Command: "help", Description: "status => чи є електрика"},
		tgbotapi.BotCommand{Command: "status", Description: "Перевірити наявність електоренергіЇ"}))

	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go processCommand(bot, update.Message.Command(), update)
		}
	}
}
