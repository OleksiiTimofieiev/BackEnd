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
			msg := tgbotapi.NewMessage(update.Message.Chat.ID,
				energyOutOfService[int(time.Now().Weekday())])

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
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

var energyOutOfService [7]string

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken)

	energyOutOfService[1] = "Понеділок: 6-10, 15-19\n"
	energyOutOfService[2] = "Вівторок: 0-4, 9-13, 18-22\n"
	energyOutOfService[3] = "Середа: 3-7, 12-16, 21-24\n"
	energyOutOfService[4] = "Четвер: 0-1, 6-10, 15-19\n"
	energyOutOfService[5] = "П'ятниця: 0-4, 9-13, 18-22\n"
	energyOutOfService[6] = "Суббота: 3-7, 12-16, 21-24\n"
	energyOutOfService[0] = "Неділя: 0-1, 6-10, 15-19"

	if err != nil {
		log.Panic(err)
	}

	bot.Send(tgbotapi.NewSetMyCommands(
		// tgbotapi.BotCommand{Command: "help", Description: "status => чи є електрика"},
		tgbotapi.BotCommand{Command: "status", Description: "Перевірити наявність електоренергіЇ"},
		tgbotapi.BotCommand{Command: "not_available_today", Description: "План відключень на сьогодні"},
		tgbotapi.BotCommand{Command: "not_available_week", Description: "План відключень на тиждень"}))

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go processCommand(bot, update.Message.Command(), update)
		}
	}
}
