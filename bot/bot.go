package bot

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramBot struct {
	botAPI        *tgbotapi.BotAPI
	admin_chat_id int64
	token         string
}

func NewTelegramBot(token string, admin_chat_id int64) (*TelegramBot, error) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		return nil, err
	}

	return &TelegramBot{
		botAPI:        bot,
		admin_chat_id: admin_chat_id,
	}, nil
}

func (t *TelegramBot) SendNotificationToAdmin(adminChannel chan interface{}, done chan struct{}) {
	defer close(done)
	data := <-adminChannel
	if identifier, ok := data.(string); ok {
		msg := tgbotapi.NewMessage(int64(t.admin_chat_id), fmt.Sprintf("Record not found: %s", identifier))
		_, err := t.botAPI.Send(msg)
		if err != nil {
			fmt.Println("Error sending notification to admin:", err)
		}
	}
}

func (t *TelegramBot) SendMessage(chatID int64, message string) {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := t.botAPI.Send(msg)
	if err != nil {
		log.Printf("Error sending message to chat ID %d: %v", chatID, err)
	}
}
