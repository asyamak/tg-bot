package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"tgbot/config"
)

func main() {
	cfg, err := config.NewConfig("config/config.json")
	//fmt.Println(cfg)
	if err != nil {
		log.Printf("error init configs: %v\n", err)
	}
	bot, err := tgbotapi.NewBotAPI(cfg.Token)
	if err != nil {
		log.Printf("error in init new bot api: %v\n", err)
	}
	//getting more information about requests being sent to TG BOT
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = cfg.TimeOut

	updates := bot.GetUpdatesChan(updateConfig)
	//if err != nil {
	//	log.Printf("error in updating congig %v\n", err)
	//}

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, I'm ur Telegram Bot!")
		bot.Send(msg)
	}
}
