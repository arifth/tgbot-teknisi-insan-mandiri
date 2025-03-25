package main

import (
	"arifthalhah/sigesit-bot/v2/clients"
	"arifthalhah/sigesit-bot/v2/config"
	"arifthalhah/sigesit-bot/v2/handlers"
	"arifthalhah/sigesit-bot/v2/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	l "log"
)

// source https://go-telegram-bot-api.dev/getting-started/
func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "*",
	}))
	bot := clients.Init()
	handlers.Init(bot)
	logger.NewLogger()
	l.Fatal(app.Listen(":" + config.Config("PORT")))
}
