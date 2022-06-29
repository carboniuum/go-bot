package main

import (
	"bufio"
	"errors"
	"fmt"
	actions "go-bot/bot"
	"log"
	"os"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var PATH string = "userinfo.txt"

func main() {
	initialize()

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

func initialize() {
	if _, err := os.Stat(PATH); errors.Is(err, os.ErrNotExist) {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter your telegram user id: ")
		userId, _ := reader.ReadString('\n')
		fmt.Print("Enter your telegram bot api token: ")
		token, _ := reader.ReadString('\n')

		file, fileErr := os.Create(PATH)

		if fileErr != nil {
			panic(fileErr)
		}

		defer file.Close()

		_, ioErr := file.WriteString(strings.TrimSpace(userId) + "\n" + strings.TrimSpace(token))

		if ioErr != nil {
			panic(ioErr)
		}	
	}
}