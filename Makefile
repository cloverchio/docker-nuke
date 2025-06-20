# Makefile for docker-nuke

BINARY_NAME=docker-nuke

build:
	go build -o $(BINARY_NAME)

test:
	go test ./...

run:
	./$(BINARY_NAME)

clean:
	rm -f $(BINARY_NAME)

all: build test