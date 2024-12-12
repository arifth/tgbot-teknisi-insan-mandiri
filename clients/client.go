package clients

import (
	"arifthalhah/sigesit-bot/v2/config"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func Init() *tgbotapi.BotAPI {
	bot, err := tgbotapi.NewBotAPI(config.Config("TELEGRAM_API_TOKEN"))
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = false
	return bot
}
