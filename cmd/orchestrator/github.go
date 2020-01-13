package main

import (
	"fmt"
	"io/ioutil"
	"context"
	"github.com/google/go-github/github"
)

type Github struct {
	Username string
	Password string
}

func (s *Github) StartSession() (*github.Client) {

	tp := github.BasicAuthTransport{
		Username: s.Username,
		Password: s.Password,
	}
	
	Client := github.NewClient(tp.Client())
	return Client
}

func (s *Github) GetFile(Organization, Repository, Path string) []byte {
	Context := context.Background()

	Client := s.StartSession()
	RawData, _ := Client.Repositories.DownloadContents(Context, Organization, Repository, Path, nil)

	body, _ := ioutil.ReadAll(RawData)
	return body
}

func (s *Github) UpdateFile(Organization, Repository, Path string, Data []byte) {

	Message := "Update by Bot"
	Branch := "master"

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

	_, _, err := Client.Repositories.UpdateFile(
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