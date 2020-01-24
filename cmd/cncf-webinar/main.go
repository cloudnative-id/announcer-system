package main

import (
	"os"
	"fmt"
	"gopkg.in/yaml.v2"
	"github.com/cloudnative-id/announcer-system/models"
	"github.com/cloudnative-id/announcer-system/handlers"
)

func main() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	var session = handlers.Github{Username: username, Password: password}
	
	var currentContent models.WebinarCNCFList
	currentContentTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/cncf-webinar/content.yaml")

	err := yaml.Unmarshal(currentContentTmpl, &currentContent)
	if err != nil {
		fmt.Println(err)
	}

	currentLen := len(currentContent.Content)
	var newContent models.WebinarCNCFList

	newContent = GetContentCNCF(currentContent)
	for _,v := range newContent.Content {
		currentContent.Content = append(currentContent.Content,v)
	}
	newLen := len(currentContent.Content)

	if newLen != currentLen {
		fmt.Println("push data CNCF webinar")
		Data, _ := yaml.Marshal(currentContent)
		session.UpdateFile("cloudnative-id","announcer-system","./resources/cncf-webinar/content.yaml",Data)
	} else {
		fmt.Println("No update in CNCF webinar")
	}
}
