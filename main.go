package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"

	"github.com/areviksol/backend_task/controller"
	"github.com/areviksol/backend_task/database"
	"github.com/areviksol/backend_task/eventbus"
	"github.com/areviksol/backend_task/model"
	"github.com/areviksol/backend_task/processor"
	"github.com/areviksol/backend_task/server"
	"github.com/areviksol/backend_task/bot"
	"github.com/joho/godotenv"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	if err := godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		fmt.Println("BOT_TOKEN environment variable is not set")
		return
	}

	adminChatIDStr := os.Getenv("ADMIN_CHAT_ID")
	if adminChatIDStr == "" {
		fmt.Println("ADMIN_CHAT_ID environment variable is not set")
		return
	}

	adminChatID, err := strconv.ParseInt(adminChatIDStr, 10, 64)
	if err != nil {
		fmt.Println("Error parsing ADMIN_CHAT_ID:", err)
		return
	}

	botInstance, err := bot.NewTelegramBot(botToken, adminChatID)
	if err != nil {
		fmt.Println("Error creating Telegram bot instance:", err)
		return
	}

	db, err := database.NewDatabase()
	if err != nil {
		fmt.Println("Error creating database:", err)
		return
	}
	model := model.NewModel(db)

	eventBus, err := eventbus.NewEventBus(botToken)

	if err != nil {
		fmt.Println("Error creating eventbus:", err)
		return
	}

	proc := processor.NewHTTPProcessor(model, eventBus)

	ctrl := controller.NewController(proc, eventBus, model, botInstance)
	srv := server.NewServer(ctrl)
	
	go func() {
		defer wg.Done()
		if err := srv.Run(); err != nil {
			fmt.Println("Error starting server:", err)
		}
	}()

	wg.Wait()
}
