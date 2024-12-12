package repositories

import (
	"arifthalhah/sigesit-bot/v2/database"
	"arifthalhah/sigesit-bot/v2/models"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func SetTask(update tgbotapi.Update) error {
	DB := database.Init()
	task := models.Task{
		ChatId: update.Message.Chat.ID,
		Task:   update.Message.Text,
	}

	if result := DB.Create(&task); result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteTask(taskId string) error {
	DB := database.Init()
	if result := DB.Where("id = ?", taskId).Delete(&models.Task{}); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetAllTasks(chatId int64) ([]models.Task, error) {
	DB := database.Init()
	var tasks []models.Task
	if result := DB.Where("chat_id = ?", chatId).Find(&tasks); result.Error != nil {
		return tasks, result.Error
	}
	return tasks, nil
}
