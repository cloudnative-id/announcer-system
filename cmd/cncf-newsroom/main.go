package main

import (
	"os"
	"fmt"
	"gopkg.in/yaml.v2"
)

func main() {
	username := os.Getenv("USERNAME")
	password := os.Getenv("PASSWORD")

	var session = Github{username: username, password: password}
	
	var currentContent ContentCNCF
	currentContentTmpl := session.GetFile("zufardhiyaulhaq","announcer-system","./resources/cncf-newsroom/content.yaml")
	
	err := yaml.Unmarshal(currentContentTmpl, &currentContent)
	if err != nil {
		fmt.Println(err)
	}

	currentLen := len(currentContent.Content)

	var newContent ContentCNCF

	announcementsURL := "https://www.cncf.io/newsroom/announcements/"
	announcementKind := "Announcements"

	newContent = GetContentCNCF(announcementsURL,announcementKind,currentContent)

	for _,v := range newContent.Content {
		currentContent.Content = append(currentContent.Content,v)
	}

	inTheNewsURL := "https://www.cncf.io/newsroom/in-the-news/"
	inTheNewsKind := "In The News"

	newContent = GetContentCNCF(inTheNewsURL,inTheNewsKind,currentContent)

	for _,v := range newContent.Content {
		currentContent.Content = append(currentContent.Content,v)
	}
	
	blogKind := "Blog"
	blogURL := "https://www.cncf.io/newsroom/blog/"

	newContent = GetContentCNCF(blogURL,blogKind,currentContent)

	for _,v := range newContent.Content {
		currentContent.Content = append(currentContent.Content,v)
	}

	newLen := len(currentContent.Content)

	if newLen != currentLen {
		Data, _ := yaml.Marshal(currentContent)
		session.UpdateFile("zufardhiyaulhaq","announcer-system","./resources/cncf-newsroom/content.yaml",Data)
	} else {
		fmt.Println("No update in CNCF Newsroom")
	}
}
