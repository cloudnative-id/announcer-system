package main

import (
	"os"
	"fmt"
	"time"
)

func main() {
	User := os.Getenv("USERNAME")
	Password := os.Getenv("PASSWORD")

	for ; ; {
		var Session = Github{User, Password}
		kubeweekly(Session)

		fmt.Printf("%v+\n", time.Now())
        time.Sleep(120 * time.Second)
	}
}