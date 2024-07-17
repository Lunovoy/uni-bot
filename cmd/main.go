package main

import (
	"log"

	"github.com/lunovoy/uni-bot/bot"
	"github.com/lunovoy/uni-bot/config"
)

func main() {
	cfg := config.NewConfig()

	telegramBot, err := bot.NewBot(cfg)
	if err != nil {
		log.Fatalf("Failed to start bot: %s", err)
	}

	telegramBot.Start()
}
