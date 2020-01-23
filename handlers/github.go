package handlers

import (
	"fmt"
	"log"
	"io/ioutil"
	"context"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Github struct {
	Username string
	Password string
}

func (s *Github) StartSession() (*github.Client) {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: s.Password},
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
	
	body, _ := ioutil.ReadAll(rawData)
	return body
}

func (s *Github) UpdateFile(organization, repository, path string, data []byte) {

	message := "Update by Bot"
	branch := "master"

	context := context.Background()
	client := s.StartSession()

	getOpts := &github.RepositoryContentGetOptions{Ref: "master"}

	res, _, _, err := client.Repositories.GetContents(
		context,
		organization,
		repository,
		path,
		getOpts,
	)

	if err != nil {
		fmt.Println(err)
	}

	_, _, err = client.Repositories.UpdateFile(
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

func (s *Github) GetURLFile(organization, repository, path string) string {
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

	URL := res.DownloadURL
	return *URL
}

func (s *Github) CreateFile(organization, repository, path string, data []byte) {

	message := "Create by Bot"
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

	_, _, err := client.Repositories.CreateFile(
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