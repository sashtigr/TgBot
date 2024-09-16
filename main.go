package main

import (
	"fmt"
	"github.com/mymmrac/telego"
	tu "github.com/mymmrac/telego/telegoutil"
	"os"
)

func main() {
	botToken := "7083057597:AAFJ7zE67MxGhDKxaycZSJHXcc9hYjvQjLU"

	bot, err := telego.NewBot(botToken, telego.WithDefaultDebugLogger())

	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	updates, _ := bot.UpdatesViaLongPolling(nil)

	defer bot.StopLongPolling()

	for update := range updates {
		if update.Message != nil {
			chatID := tu.ID(update.Message.Chat.ID)

			_, _ = bot.CopyMessage(tu.CopyMessage(chatID, chatID, update.Message.MessageID))
		}
	}
}
