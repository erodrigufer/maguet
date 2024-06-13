# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOVET=$(GOCMD) vet

BINARY_NAME=maguet
BINARY_UNIX=$(BINARY_NAME)_unix

.PHONY: build clean test vet

# Builds the Go binary for the current platform.
build:
	$(GOBUILD) -o ./build/$(BINARY_NAME) -v ./cmd/$(BINARY_NAME)

# Runs tests.
test:
	$(GOTEST) -v ./...

vet:
	$(GOVET) -v ./... 

# Removes any build artifacts from previous builds.
clean:
	$(GOCLEAN)
	rm -rf ./build/

install: build
	cp ./build/$(BINARY_NAME) ~/bin/

# Builds the Go binary and depends on the `build` target.
# run: build
# 	./$(BINARY_NAME)

# Cross-compiles the Go binary for Linux on an AMD64 architecture.
# build-linux:
# 	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v

