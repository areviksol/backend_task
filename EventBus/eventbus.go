package eventbus

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type EventBus struct {
	subscribers map[string][]chan interface{}
	mu          sync.RWMutex
	bot         *tgbotapi.BotAPI
}

func NewEventBus(botToken string) (*EventBus, error) {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return nil, err
	}

	return &EventBus{
		subscribers: make(map[string][]chan interface{}),
		bot:         bot,
	}, nil
}

func (eb *EventBus) Subscribe(event string) chan interface{} {
	eb.mu.Lock()
	defer eb.mu.Unlock()
	ch := make(chan interface{}, 1)
	eb.subscribers[event] = append(eb.subscribers[event], ch)
	return ch
}

func (eb *EventBus) Publish(event string, data interface{}) {
	eb.mu.RLock()
	defer eb.mu.RUnlock()
	for _, ch := range eb.subscribers[event] {
		ch <- data
	}
}

func (eb *EventBus) SendTelegramMessage(chatID int64, message string) error {
	msg := tgbotapi.NewMessage(chatID, message)
	_, err := eb.bot.Send(msg)
	return err
}
func (eb *EventBus) SubscribeAdmin() chan interface{} {
	return eb.Subscribe("record_not_found")
}
