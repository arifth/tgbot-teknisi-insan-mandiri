package handlers

import (
	"arifthalhah/sigesit-bot/v2/services"
	"arifthalhah/sigesit-bot/v2/templates"
	"arifthalhah/sigesit-bot/v2/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"strconv"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "buat_task":
		services.AppendNewTaskCommand(bot, update)
	}
}

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("apakah masuk callback", update.Message)
	cmd, taskId := utils.GetKeyValue(update.CallbackQuery.Data)
	switch {
	case cmd == "delete_task":
		services.DeleteTaskCallback(bot, update, taskId)
	}
}

func Init(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.AllowedUpdates = append(u.AllowedUpdates, "UpdateTypeChatMember")
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	for update := range updates {
		if update.Message == nil {
			continue
		}
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message.IsCommand() {
			Commands(bot, update)
		}
		if update.Message.Text != "" {
			continue
		}
		isValid, data, reason := utils.IsMatchFormat(update.Message.Text)
		if isValid {
			if err := services.AppendNewTask(bot, update, data); err != nil {
				fmt.Println("cannot insert into google sheets", err)
			}
			repliesToChan := templates.RepliesToChannel(update.SentFrom().UserName)
			message := utils.RequestToChannel(strconv.FormatInt(-1002294546372, 10), repliesToChan, "44")
			fmt.Println(message)
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, reason)
			bot.Send(msg)
		}
	}

}
