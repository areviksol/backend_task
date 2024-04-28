package bot

import (
	"log"

	"github.com/go-telegram-bot-api/telegram-bot-api"
)

// TelegramBot represents a Telegram bot
type TelegramBot struct {
	botAPI *tgbotapi.BotAPI
}

// NewTelegramBot creates a new instance of TelegramBot
func NewTelegramBot(token string) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{botAPI: bot}, nil
}

// SendMessage sends a message to the specified chat ID
func (t *TelegramBot) SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := t.botAPI.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat ID %d: %v", chatID, err)
	}
}
