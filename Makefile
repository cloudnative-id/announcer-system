GOCMD=go
GOBUILD=$(GOCMD) build
GOGET=$(GOCMD) get

ORCHESTRATOR_BINARY_NAME=orchestrator
KUBEWEEKLY_BINARY_NAME=kubeweekly
CNCF_NEWSROOM_BINARY_NAME=cncf-newsroom

orchestrator-build: 
		$(GOBUILD) -o $(ORCHESTRATOR_BINARY_NAME) -v ./cmd/orchestrator/

orchestrator-run:
		./$(ORCHESTRATOR_BINARY_NAME)

kubeweekly-build: 
		$(GOBUILD) -o $(KUBEWEEKLY_BINARY_NAME) -v ./cmd/kubeweekly/

kubeweekly-run:
		./$(KUBEWEEKLY_BINARY_NAME)

cncf-newsroom-build: 
		$(GOBUILD) -o $(CNCF_NEWSROOM_BINARY_NAME) -v ./cmd/cncf-newsroom/

cncf-newsroom-run:
		./$(CNCF_NEWSROOM_BINARY_NAME)

deps:
		$(GOGET) gopkg.in/yaml.v2
		$(GOGET) github.com/google/go-github/github
		$(GOGET) github.com/go-telegram-bot-api/telegram-bot-api
		$(GOGET) github.com/PuerkitoBio/goquery
		$(GOGET) golang.org/x/oauth2		