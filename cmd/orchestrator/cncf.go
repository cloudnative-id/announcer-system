package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func CNCFNewsRoom(session Github, telegramBot TelegramDispatcher)(){
	contentTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/cncf-newsroom/content.yaml")
	
	var pushRepository = false
	var contentList ContentCNCF
	var telegramContentList ContentCNCF


	yaml.Unmarshal(contentTmpl, &contentList)

	for i, s := range contentList.Content {
		if s.IsDelivered == false {
			telegramContentList.Content = append(telegramContentList.Content,s)
			contentList.Content[i].IsDelivered = true
			pushRepository = true
		}
	}

	fmt.Println("Send CNCF Newsroom message to Telegram")
	telegramBot.SendMsgTelegram(telegramContentList)

	if pushRepository {
		fmt.Println("Push updated Data CNCF Newsroom")

		Data, _ := yaml.Marshal(contentList)
		session.UpdateFile("cloudnative-id","announcer-system","./resources/cncf-newsroom/content.yaml",Data)
	} else {
		fmt.Println("No Updated in CNCF Newsroom")
	}
}
