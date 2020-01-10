GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get
BINARY_NAME=orchestrator

build: 
		$(GOBUILD) -o $(BINARY_NAME) -v ./pkg/orchestrator/

run:
		$(GOBUILD) -o $(BINARY_NAME) -v ./pkg/orchestrator/
		./$(BINARY_NAME)

deps:
		$(GOGET) gopkg.in/yaml.v2
		$(GOGET) github.com/google/go-github/github
		$(GOGET) github.com/go-telegram-bot-api/telegram-bot-api
		