package main

import (
	"os"
	"fmt"
	"strings"
	"gopkg.in/yaml.v2"
	"github.com/cloudnative-id/announcer-system/models"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func main() {

	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")

	var Session = handlers.Github{Username: User, Password: Password}
	ContentListTmpl := Session.GetFile("cloudnative-id","announcer-system","./resources/kubeweekly/kubeweekly.yaml")
	
	var CurrentContentList models.KubeweeklyList
	yaml.Unmarshal(ContentListTmpl, &CurrentContentList)

	var NewKubeweeklyTitle string
	var UpdateKubeweekly bool

	NewKubeweeklyTitle = strings.ReplaceAll(GetNewKubeweeklyTitle(), " #", "")
	UpdateKubeweekly = true

	for _, List := range CurrentContentList.Data{
		if List.Title == NewKubeweeklyTitle{
			UpdateKubeweekly = false
		}
	}

	if UpdateKubeweekly {
		fmt.Println("Start getting content for New Kubeweekly")
		
		var NewContent models.KubeweeklyContent
		NewContent = GetContentKubeweekly()

		var NewContentList models.KubeweeklyData
		NewContentList = GetContentListKubeweekly()

		NewContentYaml, _ := yaml.Marshal(NewContent)
		Session.CreateFile("cloudnative-id","announcer-system","./resources/kubeweekly/"+NewContentList.ContentFile,NewContentYaml)
		
		CurrentContentList.Data = append(CurrentContentList.Data,NewContentList)
		NewContentListYaml, _ := yaml.Marshal(CurrentContentList)
		Session.UpdateFile("cloudnative-id","announcer-system","./resources/kubeweekly/kubeweekly.yaml",NewContentListYaml)

	} else {
		fmt.Println("Kubeweekly not updated")
	}
}
