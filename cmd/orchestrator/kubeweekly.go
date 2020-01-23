package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/cloudnative-id/announcer-system/models"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func Kubeweekly(session handlers.Github, telegramBot TelegramDispatcher)(){
	ConfigTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/kubeweekly.yaml")
	
	var PushRepository = false
    var Config models.KubeweeklyList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.Data {
		if s.Status.IsDelivered == false {
			YamlTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/"+s.ContentFile)

			var Content models.KubeweeklyContent
			yaml.Unmarshal(YamlTmpl, &Content)
			
			fmt.Println("Send message to Telegram")
			telegramBot.SendMsgTelegram(Content)

			Config.Data[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Kubeweekly Data")

		Data, _ := yaml.Marshal(Config)
		session.UpdateFile("cloudnative-id","announcer-system","./resources/kubeweekly/kubeweekly.yaml",Data)
	} else {
		fmt.Println("No Updated in Kubeweekly")
	}
}
