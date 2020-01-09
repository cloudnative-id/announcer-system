package main

import (
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
	RawData, _ := Client.Repositories.DownloadContents(Context,Organization,Repository,Path, nil)

	body, _ := ioutil.ReadAll(RawData)
	return body
}