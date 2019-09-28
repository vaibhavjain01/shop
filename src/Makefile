# Go parameters
	GOCMD=go
	GOBUILD=$(GOCMD) build
	GOCLEAN=$(GOCMD) clean
	GOTEST=$(GOCMD) test
	GOGET=$(GOCMD) get
	BINARY_NAME=onestopshop
	BINARY_UNIX=$(BINARY_NAME)_unix

#all: test build
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../bin/$(BINARY_NAME) -v
test:
	$(GOTEST) -v ./...
clean:
	$(GOCLEAN)
	rm -f ../bin/$(BINARY_NAME)
	rm -f ../bin/$(BINARY_UNIX)
run:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ../bin/$(BINARY_NAME) -v
	../bin/$(BINARY_NAME)
deps:
	$(GOGET) github.com/urfave/cli
	$(GOGET) github.com/labstack/echo/...