package main

import (
	"arifthalhah/sigesit-bot/v2/clients"
	"arifthalhah/sigesit-bot/v2/config"
	"arifthalhah/sigesit-bot/v2/handlers"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	log.Fatal(app.Listen(":" + config.Config("PORT")))
}
