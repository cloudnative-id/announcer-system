package main

import (
	"os"
	"bytes"
	"strconv"
	"text/template"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

func KubeweeklyTelegram(Data KubeweeklyContent) {

	TelegramToken := os.Getenv("TELEGRAM_TOKEN")
	TelegramChatID, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))

	var Output bytes.Buffer

	Tpl, err := template.ParseFiles("templates/KubeweeklyTelegram.tmpl")
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
	Msg.DisableWebPagePreview = true

	Bot.Send(Msg)
}

func PosterMeetupTelegram(Data MeetupEvent, URL string) {
	TelegramToken := os.Getenv("TELEGRAM_TOKEN")
	TelegramChatID, _ := strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))

	var Output bytes.Buffer

	Tpl, err := template.ParseFiles("templates/PosterMeetupTelegram.tmpl")
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
	Msg.DisableWebPagePreview = true
	Bot.Send(Msg)

	Pic := tgbotapi.NewPhotoShare(int64(TelegramChatID), URL)
	Bot.Send(Pic)
}