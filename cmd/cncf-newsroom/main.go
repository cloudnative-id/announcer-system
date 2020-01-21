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
	currentContentTmpl := session.GetFile("cloudnative-id","announcer-system","./resources/cncf-newsroom/content.yaml")
	yaml.Unmarshal(currentContentTmpl, &currentContent)


	announcementsURL := "https://www.cncf.io/newsroom/announcements/"
	// inTheNewsURL := "https://www.cncf.io/newsroom/in-the-news/"
	// blogURL := "https://www.cncf.io/newsroom/blog/"

	announcementKind := "Announcements"
	// inTheNewsKind := "In The News"
	// blogKind := "Blog"

	var newContent ContentCNCF
	newContent = GetContentCNCF(announcementsURL,announcementKind,currentContent)
	
	Data, _ := yaml.Marshal(newContent)
	fmt.Println(Data)
}
