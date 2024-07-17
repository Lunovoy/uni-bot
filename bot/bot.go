package bot

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/lunovoy/uni-bot/config"
	"github.com/lunovoy/uni-bot/openai"

	"gopkg.in/telebot.v3"
)

// Структура для хранения экземпляра телеграм бота и клиента OpenAI
type Bot struct {
	bot      *telebot.Bot
	aiClient *openai.OpenAIClient
}

// Создание нового экземпляра бота с заданной конфигурацией
func NewBot(cfg *config.Config) (*Bot, error) {
	p := telebot.Settings{
		Token:  cfg.TelegramToken,
		Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := telebot.NewBot(p)
	if err != nil {
		return nil, err
	}

	aiClient := openai.NewOpenAIClient(cfg.OpenAIToken)

	bot := &Bot{
		bot:      b,
		aiClient: aiClient,
	}

	bot.initHandlers()

	return bot, nil
}

// Инициализация обработчика для обработки текстовых сообщений и фото
func (b *Bot) initHandlers() {
	// Обработка запроса с текстом
	b.bot.Handle(telebot.OnText, func(c telebot.Context) error {
		question := c.Text()
		ctx := context.Background()

		answer, err := b.aiClient.GetResponse(ctx, question)
		if err != nil {
			log.Println("Failed to get response from OpenAI:", err)
			return c.Send("Sorry, something went wrong while contacting OpenAI.")
		}

		return c.Send(answer)
	})

	// Обработка фото с текстом
	b.bot.Handle(telebot.OnPhoto, func(c telebot.Context) error {
		photo := c.Message().Photo
		caption := c.Message().Caption

		if photo == nil {
			return c.Send("Sorry, no photo received.")
		}

		file := &telebot.File{FileID: photo.FileID}
		photoPath := fmt.Sprintf("photos/%s.jpg", photo.FileID)
		err := b.bot.Download(file, photoPath)
		if err != nil {
			log.Println("Failed to download photo:", err)
			return c.Send("Sorry, something went wrong while downloading the photo.")
		}

		ctx := context.Background()

		answer, err := b.aiClient.GetResponse(ctx, caption)
		if err != nil {
			log.Println("Failed to get response from OpenAI:", err)
			return c.Send("Sorry, something went wrong while contacting OpenAI.")
		}

		// Удаляем сохраненное фото
		err = os.Remove(photoPath)
		if err != nil {
			log.Println("Failed to delete photo:", err)
		}

		return c.Send(answer)
	})
}

func (b *Bot) Start() {
	b.bot.Start()
}
