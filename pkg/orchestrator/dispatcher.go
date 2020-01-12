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

	Tpl, err := template.ParseFiles("templates/kubeweeklyTelegram.tmpl")
	if err != nil {
		panic(err)
	  }
	
	err = Tpl.Execute(&Output, Data)
	if err != nil {
		panic(err)
	  }
	
	Bot, err := tgbotapi.NewBotAPI(TelegramToken)
	if err != nil {
		panic(err)
	}

	Msg := tgbotapi.NewMessage(int64(TelegramChatID), Output.String())
	Msg.ParseMode = "markdown"
	Msg.DisableWebPagePreview = false

	Bot.Send(Msg)
}
