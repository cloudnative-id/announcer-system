package main

import (
	"log"
	"fmt"
	"io/ioutil"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	username string
	password string
}

func (s *Github) StartSession() (*github.Client) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: s.password},
	)
	tc := oauth2.NewClient(ctx, ts)
	
	client := github.NewClient(tc)
	return client
}

func (s *Github) GetFile(organization, repository, path string) []byte {
	context := context.Background()
	client := s.StartSession()

	rawData, err := client.Repositories.DownloadContents(context, organization, repository, path, nil)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(rawData)
	if err != nil {
		log.Fatal(err)
	}

	return body
}

func (s *Github) UpdateFile(organization, repository, path string, data []byte) {

	message := "Update by Bot"
	branch := "master"

	context := context.Background()
	client := s.StartSession()

	getOpts := &github.RepositoryContentGetOptions{Ref: "master"}

	res, _, _, _ := client.Repositories.GetContents(
		context,
		organization,
		repository,
		path,
		getOpts,
	)

	_, _, err := client.Repositories.UpdateFile(
		context,
		organization,
		repository,
		path,
		&github.RepositoryContentFileOptions{
			Message: &message,
			Content: data,
			Branch:  &branch,
			SHA: github.String(res.GetSHA()),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
}