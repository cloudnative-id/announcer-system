package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)

func Kubeweekly(Session Github)(){
	ConfigTmpl := Session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml")
	
	var PushRepository = false
    var Config KubeweeklyContentList

	yaml.Unmarshal(ConfigTmpl, &Config)

	for i, s := range Config.ContentLists {
		if s.Status.Delivered == false {

			YamlTmpl := Session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/"+s.Content)

			var Content KubeweeklyContent
			yaml.Unmarshal(YamlTmpl, &Content)
			
			fmt.Println("Send message to Telegram")
			KubeweeklyTelegram(Content)

			Config.ContentLists[i].Status.Delivered = true
			PushRepository = true
		}
	}

	if PushRepository {
		fmt.Println("Push updated Data")

		Data, _ := yaml.Marshal(Config)
		Session.UpdateFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml",Data)
	} else {
		fmt.Println("No Updated in Kubeweekly")
	}
}
