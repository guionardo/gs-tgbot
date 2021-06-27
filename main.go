package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/guionardo/gs_bot/behaviors"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1892910616:AAFp0ec83ElRTZOLm_14h8Lj0Wv8FWlCnzo")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)

	if err != nil {
		log.Printf("Error getting updates channel: %s", err)
	}

	var runner = behaviors.InitRunner()

	runner.AddBehavior(&behaviors.IpBehavior{})
	runner.AddBehavior(&behaviors.LinuxBehavior{})
	runner.AddBehavior(&behaviors.MemoryBehavior{})
	runner.AddBehavior(&behaviors.RebootBehavior{})
	runner.AddBehavior(&behaviors.ShutdownBehavior{})
	runner.AddBehavior(&behaviors.LastCommitBehavior{})
	runner.AddBehavior(&behaviors.UpdatingBehavior{})
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s\n", update.Message.From.UserName, update.Message.Text)
		var msg tgbotapi.MessageConfig
		var response = runner.Run(update.Message.Text)
		if len(response) == 0 {
			response = update.Message.Text
		}
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, response)
		msg.ReplyToMessageID = update.Message.MessageID

		message, err := bot.Send(msg)
		if err != nil {
			log.Printf("Error sending message: %s - %s", err, message.Text)
		}
	}
}
