GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

ORCHESTRATOR_BINARY_NAME=orchestrator
KUBEWEEKLY_BINARY_NAME=kubeweekly

orchestrator-build: 
		$(GOBUILD) -o $(ORCHESTRATOR_BINARY_NAME) -v ./cmd/orchestrator/

orchestrator-run:
		$(GOBUILD) -o $(ORCHESTRATOR_BINARY_NAME) -v ./cmd/orchestrator/
		./$(BINARY_NAME)

kubeweekly-build: 
		$(GOBUILD) -o $(KUBEWEEKLY_BINARY_NAME) -v ./cmd/kubeweekly/

kubeweekly-run:
		$(GOBUILD) -o $(KUBEWEEKLY_BINARY_NAME) -v ./cmd/kubeweekly/
		./$(BINARY_NAME)

deps:
		$(GOGET) gopkg.in/yaml.v2
		$(GOGET) github.com/google/go-github/github
		$(GOGET) github.com/go-telegram-bot-api/telegram-bot-api
		$(GOGET) github.com/PuerkitoBio/goquery		