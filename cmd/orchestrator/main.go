package main

import (
	"os"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func main() {
	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")
	var session = handlers.Github{Username: User, Password: Password}

	var telegramBot TelegramDispatcher
	telegramBot.StartBot()

	Kubeweekly(session, telegramBot)
	CNCFNewsRoom(session, telegramBot)
	CNCFWebinar(session, telegramBot)
	NewMeetup(session, telegramBot)
	PostMeetup(session, telegramBot)

}