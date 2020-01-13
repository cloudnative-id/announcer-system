package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func Meetup(Session Github)(){
	ConfigTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/meetup/ContentList.yaml")

	var PushRepository = false
	var Config MeetupContentList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.ContentLists {
		if s.Status.Delivered == false {

			YamlTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/meetup/"+s.Content)

			var Content MeetupContent
			yaml.Unmarshal(YamlTmpl, &Content)
			
			URL:= Session.GetURLFile("zufardhiyaulhaq","announcer-system","./resources/meetup/contents/"+Content.PicturePath)

			fmt.Println("Send message to Telegram")
			MeetupTelegram(Content, URL)

			Config.ContentLists[i].Status.Delivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Data")

		Data, _ := yaml.Marshal(Config)

		Session.UpdateFile("zufardhiyaulhaq", "announcer-system", "./resources/meetup/ContentList.yaml", Data)
	} else {
		fmt.Println("No Updated in Meetup")
	}
}