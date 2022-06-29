package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	actions "go-bot/bot"
	"log"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("5355590688:AAE8PjtaCNSugTQW3_R0CftpFpudCyB3fR0")

	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { 
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
			
			var msg tgbotapi.Chattable

				switch update.Message.Text {
					case "screenshot":
						photoFileBytes := actions.GetScreenshot()
						msg = tgbotapi.NewPhoto(int64(update.Message.Chat.ID), photoFileBytes)
					
					default:
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command")
				}

			bot.Send(msg)
		}
	}
}