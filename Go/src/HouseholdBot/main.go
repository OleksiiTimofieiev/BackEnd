package main

import (
	"log"
	"net/http"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	botToken       = "5547700494:AAHBQS1NUlsSpHS4Ob6VbQkWRjwz8ov7smw"
	otimofie       = 1462377126
	ntimofie       = 988002557
	homeChat       = -1001854649136
	msg            = "⚡️ Електропостачання присутнє ⚡️"
	msgElectricity = "⚡️ Електропостачання відновлене ⚡️"
	testURL        = "https://google.com"
)

const (
	status              string = "status"
	not_available_today        = "not_available_today"
	not_available_week         = "not_available_week"
)

func processCommand(bot *tgbotapi.BotAPI, command string, update tgbotapi.Update) {
	if len(command) >= 0 {
		if command == status {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, msg)
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		} else if command == not_available_today {
			day := int(time.Now().Weekday())
			if day != 7 {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, energyOutOfService[day])
				msg.ParseMode = tgbotapi.ModeHTML

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, energyOutOfService[0])

				if _, err := bot.Send(msg); err != nil {
					log.Panic(err)
				}
			}

		} else if command == not_available_week {
			var sb strings.Builder

			for i := 1; i < 7; i++ {
				sb.WriteString(energyOutOfService[i])
			}
			sb.WriteString(energyOutOfService[0])

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, sb.String())
			msg.ParseMode = tgbotapi.ModeHTML

			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)
			}
		}
	}
}

func electricityUpdate(bot *tgbotapi.BotAPI) {
	var messages []tgbotapi.MessageConfig

	messages = append(messages, tgbotapi.NewMessage(homeChat, msgElectricity))
	// messages = append(messages, tgbotapi.NewMessage(otimofie, msgElectricity))
	// messages = append(messages, tgbotapi.NewMessage(ntimofie, msgElectricity))

	for _, msg := range messages {
		_, err := bot.Send(msg)

		if err != nil {
			panic(err)
		}
	}
}

var energyOutOfService [7]string

func IsOnline() bool {
	timeout := time.Duration(5000 * time.Millisecond)
	client := http.Client{
		Timeout: timeout,
	}
	_, err := client.Get(testURL)

	if err != nil {
		return false
	}

	return true
}

func main() {
	for IsOnline() == false {
		time.Sleep(time.Second)
	}

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}

	energyOutOfService[1] = "Понеділок: <u>0</u>, 3-5, <u>6-9</u>, 12-14, <u>15-18</u>, 21-23\n"
	energyOutOfService[2] = "Вівторок: <u>0-3</u>, 6-8, <u>9-12</u>, 15-17, <u>18-21</u>\n"
	energyOutOfService[3] = "Середа: 0-2, <u>3-6</u>, 9-11, <u>12-15</u>, 18-20, <u>21-23</u>\n"
	energyOutOfService[4] = "Четвер: <u>0</u>, 3-5, <u>6-9</u>, 12-14, <u>15-18</u>, 21-23\n"
	energyOutOfService[5] = "П'ятниця: <u>0-3</u>, 6-8, <u>9-12</u>, 15-17, <u>18-21</u>\n"
	energyOutOfService[6] = "Суббота: 0-2, <u>3-6</u>, 9-11, <u>12-15</u>, 18-20, <u>21-23</u>\n"
	energyOutOfService[0] = "Неділя: <u>0</u>, 3-5, <u>6-9</u>, 12-14, <u>15-18</u>, 21-23"

	bot.Send(tgbotapi.NewSetMyCommands(
		tgbotapi.BotCommand{Command: status, Description: "Перевірити наявність електоренергіЇ"},
		tgbotapi.BotCommand{Command: not_available_today, Description: "План відключень на сьогодні"},
		tgbotapi.BotCommand{Command: not_available_week, Description: "План відключень на тиждень"}),
	)

	bot.Debug = false

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	electricityUpdate(bot)

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {
			go processCommand(bot, update.Message.Command(), update)
		}
	}
}
