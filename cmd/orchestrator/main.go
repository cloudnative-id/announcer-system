package main

import (
	"os"
)

func main() {
	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")

	var Session = Github{User, Password}

	Kubeweekly(Session)
	PosterMeetup(Session)
	PostMeetup(Session)

}