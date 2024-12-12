package handlers

import (
	"arifthalhah/sigesit-bot/v2/services"
	"arifthalhah/sigesit-bot/v2/utils"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func Commands(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	switch update.Message.Command() {
	case "start":
		services.Start(bot, update)
	case "buat_task":
		services.AppendNewTaskCommand(bot, update)
		//case "set_todo":
		//	services.SetTask(bot, update)
		//case "delete_todo":
		//	services.DeleteTask(bot, update)
		//case "show_all_todos":
		//	services.ShowAllTasks(bot, update)
	}
}

func Callbacks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	fmt.Println("apakah masuk callback", update.Message)
	//cmd, taskId := utils.GetKeyValue(update.CallbackQuery.Data)
	//switch {
	//case cmd == "delete_task":
	//	services.DeleteTaskCallback(bot, update, taskId)
	//}
}

func Init(bot *tgbotapi.BotAPI) {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)
	// Loop through each update.
	for update := range updates {
		if update.Message == nil {
			continue
		}
		fmt.Println(update.Message)
		//fmt.Println(update.Message.Document.FileName, "\n", update.Message.Document.FileID)
		isMatch := utils.IsMatchFormat(update.Message.Text)
		if isMatch {
			services.AppendNewTask(bot, update)
		}
		if update.CallbackQuery != nil {
			Callbacks(bot, update)
		} else if update.Message.IsCommand() {
			Commands(bot, update)
		}
		//else {
		//	Messages(bot, update)
		//}
	}
}
