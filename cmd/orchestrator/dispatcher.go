package main

import (
	"os"
	"bytes"
	"strconv"
	"text/template"
	"github.com/go-telegram-bot-api/telegram-bot-api"
)

type TelegramDispatcher struct {
	Token string
	ChatID int
	Bot *tgbotapi.BotAPI
	Msg tgbotapi.MessageConfig
	Pic tgbotapi.PhotoConfig
}

func (t *TelegramDispatcher) GetCredential() {
	var err error

	t.Token = os.Getenv("TELEGRAM_TOKEN")
	t.ChatID, err = strconv.Atoi(os.Getenv("TELEGRAM_CHATID"))
	if err != nil {
		panic(err)
	  }
}

func (t *TelegramDispatcher) StartBot() {
	var err error

	t.GetCredential()
	t.Bot, err = tgbotapi.NewBotAPI(t.Token)
	if err != nil {
		panic(err)
	  }
}

func (t *TelegramDispatcher) MessageBot(output bytes.Buffer) {
	t.Msg = tgbotapi.NewMessage(int64(t.ChatID), output.String())
	t.Msg.ParseMode = "markdown"
	t.Msg.DisableWebPagePreview = true
}

func (t *TelegramDispatcher) PictureBot(url string) {
	t.Pic = tgbotapi.NewPhotoShare(int64(t.ChatID), url)
}

func (t *TelegramDispatcher) SendMsgTelegram(arg interface{}) {
	var tmplFile string
	var output bytes.Buffer

	switch arg.(type) {
    case KubeweeklyContent:
        tmplFile = "templates/KubeweeklyTelegram.tmpl"
    case ContentCNCF:
		tmplFile = "templates/CNCFTelegram.tmpl"
	case MeetupEvent:
		tmplFile = "templates/NewMeetupTelegram.tmpl"
	case PostMeetupEventList:
        tmplFile = "templates/PostMeetupTelegram.tmpl"
	}
	
	tpl, err := template.ParseFiles(tmplFile)
	if err != nil {
		panic(err)
	  }
	
	err = tpl.Execute(&output, arg)
	if err != nil {
		panic(err)
	  }
	
	t.MessageBot(output)
	t.Bot.Send(t.Msg)
}


func (t *TelegramDispatcher) SendPicTelegram(url string) {
	t.PictureBot(url)
	t.Bot.Send(t.Pic)
}


