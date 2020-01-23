package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func NewMeetup(session Github, telegramBot TelegramDispatcher)(){
	ConfigTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/meetup/poster/EventList.yaml")

	var PushRepository = false
	var Config MeetupEventList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.EventLists {
		if s.Status.IsDelivered == false {
			YamlTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/meetup/poster/"+s.Event)

			var Event MeetupEvent
			yaml.Unmarshal(YamlTmpl, &Event)
			
			URL:= session.GetURLFile("cloudnative-id","announcer-system","./resources/meetup/poster/event/"+Event.PicturePath)

			fmt.Println("Send message to Telegram")
			telegramBot.SendMsgTelegram(Event)
			telegramBot.SendPicTelegram(URL)

			Config.EventLists[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Data")

		Data, _ := yaml.Marshal(Config)
		session.UpdateFile("cloudnative-id", "announcer-system", "./resources/meetup/poster/EventList.yaml", Data)
	} else {
		fmt.Println("No Updated in Meetup")
	}
}

func PostMeetup(session Github, telegramBot TelegramDispatcher)(){
	ConfigTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/meetup/postevent/EventList.yaml")

	var PushRepository = false
	var Config PostMeetupEvent

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.EventLists {
		if s.Status.IsDelivered == false {
			URL:= session.GetURLFile("cloudnative-id","announcer-system","./resources/meetup/postevent/"+s.PicturePath)

			fmt.Println("Send message to Telegram")
			telegramBot.SendMsgTelegram(s)
			telegramBot.SendPicTelegram(URL)

			Config.EventLists[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Data")

		Data, _ := yaml.Marshal(Config)
		session.UpdateFile("cloudnative-id", "announcer-system", "./resources/meetup/postevent/EventList.yaml", Data)
	} else {
		fmt.Println("No Updated in Meetup")
	}
}
