package main

import (
	"os"
	"fmt"
	"gopkg.in/yaml.v2"
)

func main() {

	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")

	var Session = Github{User, Password}
	ContentListTmpl := Session.GetFile("zufardhiyaulhaq","announcer-system","./resources/kubeweekly/ContentList.yaml")
	
	var CurrentContentList KubeweeklyContentList
	yaml.Unmarshal(ContentListTmpl, &ContentListTmpl)

	var NewKubeweeklyTitle string
	var UpdateKubeweekly bool

	NewKubeweeklyTitle = GetNewKubeweeklyTitle()
	UpdateKubeweekly = true

	for _, List := range CurrentContentList.ContentLists{
		if List.Title == NewKubeweeklyTitle{
			UpdateKubeweekly = false
		}
	}

	if UpdateKubeweekly {
		fmt.Println("Start getting content for New Kubeweekly")
		
		var NewContent KubeweeklyContent
		NewContent = GetContentKubeweekly()
		fmt.Println(NewContent)

		var NewContentList ContentList
		NewContentList = GetContentListKubeweekly()
		fmt.Println(NewContentList)
		
	} else {
		fmt.Println("Kubeweekly not updated")
	}
}
