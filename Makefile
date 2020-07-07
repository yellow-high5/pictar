GOCMD=go
GORUN=$(GOCMD) run
GOBUILD=$(GOCMD) build
GOINSTALL=$(GOCMD) install
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVERSION=$(GOCMD) version

BINARY_NAME=pictar

run: 
	$(GORUN) main.go

build:
	$(GOBUILD) -o $(BINARY_NAME) -v

install:
	$(GOINSTALL) github.com/yellow-high5/pictar

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

version:
	$(GOVERSION)
