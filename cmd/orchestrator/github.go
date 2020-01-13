package main

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
	
	Client := github.NewClient(tc)
	return Client
}

func (s *Github) GetFile(Organization, Repository, Path string) []byte {
	Context := context.Background()
	Client := s.StartSession()
	
	RawData, err := Client.Repositories.DownloadContents(Context, Organization, Repository, Path, nil)
	if err != nil {
		log.Fatal(err)
	}
	
	body, _ := ioutil.ReadAll(RawData)
	return body
}

func (s *Github) UpdateFile(Organization, Repository, Path string, Data []byte) {

	Message := "Update by Bot"
	Branch := "master"

	Context := context.Background()
	Client := s.StartSession()

	getOpts := &github.RepositoryContentGetOptions{Ref: "master"}

	res, _, _, err := Client.Repositories.GetContents(
		Context,
		Organization,
		Repository,
		Path,
		getOpts,
	)

	if err != nil {
		fmt.Println(err)
	}

	_, _, err = Client.Repositories.UpdateFile(
		Context,
		Organization,
		Repository,
		Path,
		&github.RepositoryContentFileOptions{
			Message: &Message,
			Content: Data,
			Branch:  &Branch,
			SHA: github.String(res.GetSHA()),
		},
	)

	if err != nil {
		fmt.Println(err)
	}
}

func (s *Github) GetURLFile(Organization, Repository, Path string) string {
	Context := context.Background()
	Client := s.StartSession()

	getOpts := &github.RepositoryContentGetOptions{Ref: "master"}

	res, _, _, _ := Client.Repositories.GetContents(
		Context,
		Organization,
		Repository,
		Path,
		getOpts,
	)

	URL := res.DownloadURL
	return *URL
}