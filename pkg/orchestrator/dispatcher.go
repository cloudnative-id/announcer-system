package main

import (
	"os"
	"bytes"
	"strconv"
	"text/template"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func SendTelegram(Data KubeweeklyContent) {

	TelegramToken := os.Getenv("TELEGRAM_TOKEN")
	TelegramChatID, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))

	var Output bytes.Buffer

	tpl, err := template.ParseFiles("templates/kubeweeklyTelegram.tmpl")
	if err != nil {
		panic(err)
	  }
	
	err = tpl.Execute(&Output, Data)
	if err != nil {
		panic(err)
	  }
	
	bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		panic(err)
	}

	msg := tgbotapi.NewMessage(int64(TelegramChatID), Output.String())
	msg.ParseMode = "markdown"
	bot.Send(msg)
}
