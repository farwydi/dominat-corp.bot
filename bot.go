package main

import (
	"log"
	"gopkg.in/telegram-bot-api.v4"
	"os"
	"net/http"
	"fmt"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))

	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)

	// Setting web hook https protocol
	hook := tgbotapi.NewWebhook("https://" + os.Getenv("PROGRAM_NAME") + ".herokuapp.com/" + bot.Token)
	_, err = bot.SetWebhook(hook)

	if err != nil {
		fmt.Errorf("Problem in setting Webhook: " + err.Error())
	}

	// Set Handler for http server
	updates := bot.ListenForWebhook("/" + bot.Token)

	// Start http server on PORT
	go http.ListenAndServe(":"+os.Getenv("PORT"), nil)

	// main loop get updates
	for update := range updates {
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text + " bot read")

		// Echo
		bot.Send(msg)
	}
}
