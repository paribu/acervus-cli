.PHONY: all windows linux darwin clean

BINARY_NAME=acervus
VERSION=$(shell git describe --tags --always --dirty)

all: windows linux darwin

windows:
	@GOOS=windows GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)-$(VERSION)-windows-amd64

linux:
	@GOOS=linux GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)-$(VERSION)-linux-amd64

darwin:
	@GOOS=darwin GOARCH=amd64 go build -o ./bin/$(BINARY_NAME)-$(VERSION)-darwin-amd64

clean:
	@rm -f ./bin/$(BINARY_NAME)-*-amd64

dev:
	go build -o bin/acervus main.go
