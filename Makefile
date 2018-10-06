SHELL=/bin/bash
OUTPUT:=payments
GOOS:=linux
GOARCH = amd64
ENTRYPOINT:=./main.go
BUILD_DIR?=build

ifeq ($(OS),Windows_NT)
	GOOS := windows
	EXT := .exe
else
	UNAME_S := $(shell uname -s)
	ifeq ($(UNAME_S),Linux)
		GOOS := linux
	endif
	ifeq ($(UNAME_S),Darwin)
		GOOS := darwin
	endif
endif

.PHONY: all clean build test fmt vet run

default: all

all: clean fmt vet test build

clean:
	rm -rf $(BUILD_DIR)/*

fmt:
	go fmt ./...

vet:
	go vet ./...

build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=${GOARCH} go build -o ./$(BUILD_DIR)/$(OUTPUT)$(EXT) $(ENTRYPOINT)

test:
	go test ./...

testcover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

run:
	./$(BUILD_DIR)/$(OUTPUT)$(EXT) serve
