package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func PosterMeetup(Session Github)(){
	ConfigTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/meetup/poster/EventList.yaml")

	var PushRepository = false
	var Config MeetupEventList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.EventLists {
		if s.Status.IsDelivered == false {

			YamlTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/meetup/poster/"+s.Event)

			var Event MeetupEvent
			yaml.Unmarshal(YamlTmpl, &Event)
			
			URL:= Session.GetURLFile("zufardhiyaulhaq","announcer-system","./resources/meetup/poster/event/"+Event.PicturePath)

			fmt.Println("Send message to Telegram")
			PosterMeetupTelegram(Event, URL)

			Config.EventLists[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Data")

		Data, _ := yaml.Marshal(Config)

		Session.UpdateFile("zufardhiyaulhaq", "announcer-system", "./resources/meetup/poster/EventList.yaml", Data)
	} else {
		fmt.Println("No Updated in Meetup")
	}
}

func PostMeetup(Session Github)(){
	ConfigTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/meetup/postevent/EventList.yaml")

	// var PushRepository = false
	var Config PostMeetupEvent

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.EventLists {
		if s.Status.IsDelivered == false {

			URL:= Session.GetURLFile("zufardhiyaulhaq","announcer-system","./resources/meetup/poster/eventevent/"+s.PicturePath)

			fmt.Println("Send message to Telegram")
			PostMeetupTelegram(s, URL)

			Config.EventLists[i].Status.IsDelivered = true
			// PushRepository = true
		}
	}

	// if PushRepository {
	// 	fmt.Println("Push updated Data")

	// 	Data, _ := yaml.Marshal(Config)

	// 	Session.UpdateFile("zufardhiyaulhaq", "announcer-system", "./resources/meetup/postevent/EventList.yaml", Data)
	// } else {
	// 	fmt.Println("No Updated in Meetup")
	// }
}
