package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/cloudnative-id/announcer-system/models"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func NewMeetup(session handlers.Github, telegramBot TelegramDispatcher)(){
	ConfigTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/new-meetup/newMeetup.yaml")

	var PushRepository = false
	var Config models.NewMeetupList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.Data {
		if s.Status.IsDelivered == false {
			YamlTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/new-meetup/"+s.ContentFile)

			var Event models.NewMeetupContent
			yaml.Unmarshal(YamlTmpl, &Event)
			
			URL:= session.GetURLFile("cloudnative-id","announcer-system","./resources/new-meetup/meetup/"+Event.PictureFile)

			fmt.Println("Send New Meetup message to Telegram")
			telegramBot.SendMsgTelegram(Event)
			telegramBot.SendPicTelegram(URL)

			Config.Data[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated New Meetup Data")

		Data, _ := yaml.Marshal(Config)
		session.UpdateFile("cloudnative-id", "announcer-system", "./resources/new-meetup/newMeetup.yaml", Data)
	} else {
		fmt.Println("No Updated in New Meetup")
	}
}

func PostMeetup(session handlers.Github, telegramBot TelegramDispatcher)(){
	ConfigTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/post-meetup/postMeetup.yaml")

	var PushRepository = false
	var Config models.PostMeetupList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.Content {
		if s.Status.IsDelivered == false {
			URL:= session.GetURLFile("cloudnative-id","announcer-system","./resources/post-meetup/"+s.PictureFile)

			fmt.Println("Send Post Meetup message to Telegram")
			telegramBot.SendMsgTelegram(s)
			telegramBot.SendPicTelegram(URL)

			Config.Content[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Post Meetup Data")

		Data, _ := yaml.Marshal(Config)
		session.UpdateFile("cloudnative-id", "announcer-system", "./resources/post-meetup/postMeetup.yaml", Data)
	} else {
		fmt.Println("No Updated in Post Meetup")
	}
}
