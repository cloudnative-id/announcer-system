package main

import (
	"os"
	"fmt"
)

func main() {
	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")

	var Session = Github{User, Password}
	Data := Session.GetFile("cloudnative-id","kubeweekly-scrapper","contentList.yaml")
	fmt.Println(Data)
}