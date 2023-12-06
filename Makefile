.PHONY: all

build:
	go build -o bin/acervus main.go

all: build