package test

import (
	"testing"

	"github.com/lunovoy/uni-bot/bot"
	"github.com/lunovoy/uni-bot/config"
)

func TestNewBot(t *testing.T) {
	cfg := config.NewConfig()
	_, err := bot.NewBot(cfg)
	if err != nil {
		t.Fatalf("Failed to create bot: %s", err)
	}
}
