package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func Kubeweekly(session Github, telegramBot TelegramDispatcher)(){
	ConfigTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml")
	
	var PushRepository = false
    var Config KubeweeklyContentList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.ContentLists {
		if s.Status.IsDelivered == false {
			YamlTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/"+s.Content)

			var Content KubeweeklyContent
			yaml.Unmarshal(YamlTmpl, &Content)
			
			fmt.Println("Send message to Telegram")
			telegramBot.SendMsgTelegram(Content)

			Config.ContentLists[i].Status.IsDelivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Data")

		Data, _ := yaml.Marshal(Config)
		session.UpdateFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml",Data)
	} else {
		fmt.Println("No Updated in Kubeweekly")
	}
}
