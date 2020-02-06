package main

import (
	"os"
	"sync"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func main() {
	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")
	var session = handlers.Github{Username: User, Password: Password}

	var telegramBot TelegramDispatcher
	telegramBot.StartBot()

	total := 3
    var wg sync.WaitGroup
    wg.Add(total)

	go Kubeweekly(session, telegramBot, &wg)
	go CNCFNewsRoom(session, telegramBot, &wg)
	go CNCFWebinar(session, telegramBot, &wg)
	wg.Wait()

	NewMeetup(session, telegramBot)
	PostMeetup(session, telegramBot)
}