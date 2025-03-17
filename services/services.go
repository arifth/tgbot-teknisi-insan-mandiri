package services

import (
	"arifthalhah/sigesit-bot/v2/config"
	"arifthalhah/sigesit-bot/v2/keyboards"
	"arifthalhah/sigesit-bot/v2/repositories"
	"arifthalhah/sigesit-bot/v2/repositories/Task"
	"arifthalhah/sigesit-bot/v2/templates"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func AppendNewTask(bot *tgbotapi.BotAPI, update tgbotapi.Update, data []string) error {
	//variable initiation
	dateTime := update.Message.Time()
	parsedTime := dateTime.Format("02/01/2006")
	hourMinute := dateTime.Format("15:04 MST")
	userName := fmt.Sprintf("%s %s", update.Message.From.FirstName, update.Message.From.LastName)
	text := templates.RepliesSuccesInsertDataToSheet(config.Config("SPREADSHEET_ID"))

	//initiate Sheets service
	srv := repositories.Init()

	//Add data from defined user and jam
	data = append([]string{parsedTime, hourMinute, userName}, data...)
	response, err := repositories.InsertIntoSheet(srv, config.Config("SPREADSHEET_ID"), config.Config("SHEETS_ID"), "A2:N2", data)
	if err != nil {
		log.Fatal("cannot insert into sheet", err)
	}
	if response.HTTPStatusCode == 200 {
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
		if _, err := bot.Send(msg); err != nil {
			return err
		}
	}

	sheetID := repositories.GetSheetID(srv, config.Config("SPREADSHEET_ID"), config.Config("SHEETS_ID"))
	fmt.Println(sheetID)
	return nil
}

func StartCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := templates.RepliesToCreateNewTask()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func AppendNewTaskCommand(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	//TODO: add validation for user before inserting data to sheet
	text := templates.RepliesToCreateNewTask()

	parsedData := update.Message.Chat
	fmt.Println(parsedData)

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)

	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func CreateNewTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := templates.RepliesSuccess()
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboards.CmdKeyboard()
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Please, write todo."
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func SetTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Todo successfully created"

	err := Task.SetTask(update)
	if err != nil {
		text = "Couldnt set task"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteTask(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	data, _ := Task.GetAllTasks(update.Message.Chat.ID)
	var btns []tgbotapi.InlineKeyboardButton
	for i := 0; i < len(data); i++ {
		btn := tgbotapi.NewInlineKeyboardButtonData(data[i].Task, "delete_task="+data[i].ID.String())
		btns = append(btns, btn)
	}

	var rows [][]tgbotapi.InlineKeyboardButton
	for i := 0; i < len(btns); i += 2 {
		if i < len(btns) && i+1 < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i], btns[i+1])
			rows = append(rows, row)
		} else if i < len(btns) {
			row := tgbotapi.NewInlineKeyboardRow(btns[i])
			rows = append(rows, row)
		}
	}
	fmt.Println(len(rows))
	var keyboard = tgbotapi.InlineKeyboardMarkup{InlineKeyboard: rows}
	//keyboard.InlineKeyboard = rows

	text := "Please, select todo you want to delete"
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyMarkup = keyboard
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func DeleteTaskCallback(bot *tgbotapi.BotAPI, update tgbotapi.Update, taskId string) {
	text := "Task successfully deleted"

	err := Task.DeleteTask(taskId)
	if err != nil {
		text = "Couldnt delete task"
	}

	msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}

func ShowAllTasks(bot *tgbotapi.BotAPI, update tgbotapi.Update) {
	text := "Tasks: \n"

	tasks, err := Task.GetAllTasks(update.Message.Chat.ID)
	if err != nil {
		text = "Couldnt get tasks"
	}

	for i := 0; i < len(tasks); i++ {
		text += tasks[i].Task + " \n"
	}

	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	if _, err := bot.Send(msg); err != nil {
		panic(err)
	}
}
