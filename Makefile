# Go parameters
GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOLINT := ~/go/bin/golangci-lint

# Directories
SRC_DIR := $(shell pwd)/cmd/urlShortner
BIN_DIR := $(shell pwd)/bin

# Targets
BINARY_NAME := url-shortner

.PHONY: all build clean test lint

all: build

build:
	$(GOBUILD) -o $(BIN_DIR)/$(BINARY_NAME) $(SRC_DIR)

clean:
	$(GOCLEAN)
	rm -rf $(BIN_DIR)

test:
	$(GOTEST) -v $(SRC_DIR)/...

lint:
	$(GOLINT) run $(shell pwd)/...


