package main

import (
	"fmt"
	"sync"
	"gopkg.in/yaml.v2"
	"github.com/cloudnative-id/announcer-system/models"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func CNCFNewsRoom(session handlers.Github, telegramBot TelegramDispatcher, wg *sync.WaitGroup)(){
	defer wg.Done()
	contentTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/cncf-newsroom/content.yaml")
	
	var pushRepository = false
	var contentList models.NewsroomCNCFList
	var telegramContentList models.NewsroomCNCFList

	yaml.Unmarshal(contentTmpl, &contentList)

	for i, s := range contentList.Content {
		if s.IsDelivered == false {
			telegramContentList.Content = append(telegramContentList.Content,s)
			contentList.Content[i].IsDelivered = true
			pushRepository = true
		}
	}

	if pushRepository {
		fmt.Println("Send CNCF Newsroom message to Telegram")
		telegramBot.SendMsgTelegram(telegramContentList)

		fmt.Println("Push updated Data CNCF Newsroom")
		Data, _ := yaml.Marshal(contentList)
		session.UpdateFile("cloudnative-id","announcer-system","./resources/cncf-newsroom/content.yaml",Data)
	} else {
		fmt.Println("No Updated in CNCF Newsroom")
	}
}

func CNCFWebinar(session handlers.Github, telegramBot TelegramDispatcher, wg *sync.WaitGroup)(){
	defer wg.Done()
	contentTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/cncf-webinar/content.yaml")
	
	var pushRepository = false
	var contentList models.WebinarCNCFList
	var telegramContentList models.WebinarCNCFList

	yaml.Unmarshal(contentTmpl, &contentList)

	for i, s := range contentList.Content {
		if s.IsDelivered == false {
			telegramContentList.Content = append(telegramContentList.Content,s)
			contentList.Content[i].IsDelivered = true
			pushRepository = true
		}
	}

	if pushRepository {
		fmt.Println("Send CNCF Webinar message to Telegram")
		telegramBot.SendMsgTelegram(telegramContentList)

		fmt.Println("Push updated Data CNCF Webinar")
		Data, _ := yaml.Marshal(contentList)
		session.UpdateFile("cloudnative-id","announcer-system","./resources/cncf-webinar/content.yaml",Data)
	} else {
		fmt.Println("No Updated in CNCF Webinar")
	}
}
