package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	actions "go-bot/bot"
	base "go-bot/settings"
	"log"
	"os"
)

var CmdType string = ""

func main() {
	base.Initialize()

	bot, err := tgbotapi.NewBotAPI(base.TOKEN)

	if err != nil {
		log.Panic("Something wrong with your credentials. Please, restart the executable on your machine and type credentials instead of press any key")
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { 
			var msg tgbotapi.Chattable

				if (update.Message.Chat.ID != base.USERID) {
					bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Something went wrong. Restart the executable on your machine and type credentials instead of press any key"))
					os.Exit(1)
				}

				switch update.Message.Text {
					case "/start":
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Welcome to this bot! See what it can do...\n/help")

					case "/help":
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, base.COMMANDS)

					case "/screenshot":
						msg = tgbotapi.NewPhoto(int64(update.Message.Chat.ID), actions.GetScreenshot())

					case "/kill_proc":
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Which one?"))
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, actions.GetAllProcesses()))
						CmdType = "/kill_proc"
						continue
					
					case "/browse":
						bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Url..."))
						CmdType = "/browse"
						continue
					
					default:
						msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Unknown command")
				}

				if len(CmdType) > 0 {
					switch CmdType {
						case "/kill_proc":
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, actions.KillProcess(update.Message.Text))
							CmdType = ""
						case "/browse":
							msg = tgbotapi.NewMessage(update.Message.Chat.ID, "Opened")
							actions.OpenBrowser(update.Message.Text)
							CmdType = ""
					}
				}

			bot.Send(msg)
		}
	}
}