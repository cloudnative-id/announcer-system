package main

import (
	"bytes"
	"context"
	"github.com/google/go-github/github"
)
type Github struct {
	Username string
	Password string
}

type RepositoryContentGetOptions struct {
    Ref string `url:"ref,omitempty"`
}

func (s *Github) StartSession() (*github.Client) {

	tp := github.BasicAuthTransport{
		Username: s.Username,
		Password: s.Password,
	}
	
	Client := github.NewClient(tp.Client())
	return Client
}

func (s *Github) GetFile(Organization, Repository, Path string) string {
	Context := context.Background()

	Client := s.StartSession()
	RawData, _ := Client.Repositories.DownloadContents(Context,Organization,Repository,Path, nil)

	buf := new(bytes.Buffer)
    buf.ReadFrom(RawData)
	return (buf.String())
}