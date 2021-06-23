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

	}

	var runner = behaviors.InitRunner()

	runner.AddBehavior("ip", &behaviors.IpBehavior{})
	runner.AddBehavior("linux", &behaviors.LinuxBehavior{})
	runner.AddBehavior("mem", &behaviors.MemoryBehavior{})
	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
		var msg tgbotapi.MessageConfig
		var response = runner.Run(update.Message.Text)
		if len(response) == 0 {
			response = update.Message.Text
		}
		msg = tgbotapi.NewMessage(update.Message.Chat.ID, response)
		// if update.Message.Text == "ip" {
		// 	msg = ResponseIP(update.Message.Chat.ID)
		// } else if update.Message.Text == "mem" {
		// 	msg = ResponseMem(update.Message.Chat.ID)
		// } else {
		// 	msg = tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)
		// }

		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
