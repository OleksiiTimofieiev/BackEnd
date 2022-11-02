package main

import (
	// "fmt"

	"log"
	"strings"
	"time"

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
		} else if command == "not_available_today" {
			day := int(time.Now().Weekday())
			if day != 7 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					energyOutOfService[day])

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					energyOutOfService[0])

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}

		} else if command == "not_available_week" {
			var sb strings.Builder

			for i := 1; i < 7; i++ {
				sb.WriteString(energyOutOfService[i])
			}
			sb.WriteString(energyOutOfService[0])

			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				sb.String())

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

// find device on network: ssh orangepi@172.20.10.3

var energyOutOfService [7]string

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	energyOutOfService[1] = "Понеділок: 6-9, 15-18\n"
	energyOutOfService[2] = "Вівторок: 0-3, 9-12, 18-21\n"
	energyOutOfService[3] = "Середа: 3-6, 12-15, 21-23\n"
	energyOutOfService[4] = "Четвер: 0-1, 6-9, 15-18\n"
	energyOutOfService[5] = "П'ятниця: 0-3, 9-12, 18-21\n"
	energyOutOfService[6] = "Суббота: 3-6, 12-15, 21-23\n"
	energyOutOfService[0] = "Неділя: 0-1, 6-9, 15-18"

	bot.Send(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: "status", Description: "Перевірити наявність електоренергіЇ"},
		tgbotapi.BotCommand{Command: "not_available_today", Description: "План відключень на сьогодні"},
		tgbotapi.BotCommand{Command: "not_available_week", Description: "План відключень на тиждень"}))

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
