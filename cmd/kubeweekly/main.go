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
	ContentListTmpl := Session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml")
	
	var CurrentContentList KubeweeklyContentList
	yaml.Unmarshal(ContentListTmpl, &CurrentContentList)

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

		var NewContentList ContentList
		NewContentList = GetContentListKubeweekly()

		NewContentYaml, _ := yaml.Marshal(NewContent)
		Session.CreateFile("cloudnative-id","announcer-system","./resources/kubeweekly/"+NewContentList.Content,NewContentYaml)
		
		CurrentContentList.ContentLists = append(CurrentContentList.ContentLists,NewContentList)
		NewContentListYaml, _ := yaml.Marshal(CurrentContentList)
		Session.UpdateFile("cloudnative-id","announcer-system","./resources/kubeweekly/ContentList.yaml",NewContentListYaml)

	} else {
		fmt.Println("Kubeweekly not updated")
	}
}
