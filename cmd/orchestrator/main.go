package main

import (
	"os"
)

func main() {
	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")
	var session = Github{User, Password}

	var telegramBot TelegramDispatcher
	telegramBot.StartBot()

	Kubeweekly(session, telegramBot)
	CNCFNewsRoom(session, telegramBot)
	NewMeetup(session, telegramBot)
	PostMeetup(session, telegramBot)

}