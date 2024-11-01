package main

import (
	"fmt"
	tbapi "github.com/OvyFlash/telegram-bot-api/v6"
	"github.com/pkarpovich/turtle-hub/services/telegram-bot/bot"
	"github.com/pkarpovich/turtle-hub/services/telegram-bot/config"
	"github.com/pkarpovich/turtle-hub/services/telegram-bot/events"
	"log"
)

func main() {
	log.Printf("[INFO] Starting app")

	cfg, err := config.Init()
	if err != nil {
		log.Fatalf("[ERROR] Error reading config: %s", err)
	}

	if err := run(cfg); err != nil {
		log.Fatalf("[ERROR] Error running app: %s", err)
	}
}

func run(cfg *config.Config) error {
	tbAPI, err := tbapi.NewBotAPI(cfg.Telegram.Token)
	if err != nil {
		return fmt.Errorf("failed to create Telegram events: %w", err)
	}

	tgListener := &events.TelegramListener{
		SuperUsers: cfg.Telegram.SuperUsers,
		TbAPI:      tbAPI,
		Bot: bot.MultiBot{
			bot.NewOpenai(),
		},
	}

	if err := tgListener.Do(); err != nil {
		return fmt.Errorf("failed to start Telegram listener: %w", err)
	}

	return nil
}
