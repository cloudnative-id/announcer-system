package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
)
func kubeweekly(Session Github)(){
	Data := Session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml")

    var config KubeweeklyContentList
	yaml.Unmarshal(Data, &config)
	
	fmt.Println(config)
}
