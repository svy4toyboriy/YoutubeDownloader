package main

import (
	"botyra/internal/services"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	bot, err := tgbotapi.NewBotAPI("KEY")
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	u := tgbotapi.NewUpdate(0)
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		services.HandleUpdate(bot, update)
	}
	return nil
}
