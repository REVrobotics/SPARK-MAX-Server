# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

CURRENT_DIR=$(shell pwd)

export CGO_CFLAGS=-IC:/Users/Will/go/src/github.com/willtoth/USB-BLDC-TOOL/ -g -O2
export CGO_LDFLAGS=-LC:/Users/Will/go/src/github.com/willtoth/USB-BLDC-TOOL/ -g -O2

$(info ${CGO_CFLAGS})

# Binary names
BINARY_NAME_WINDOWS=SparkMax.exe
BINARY_UNIX=SparkMax.out

ifeq ($(OS),Windows_NT)
	BINARY_NAME=${BINARY_NAME_WINDOWS}
else
	BINARY_NAME=${BINARY_UNIX}
endif

all: build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/pebbe/zmq4
	$(GOGET) -u github.com/spf13/cobra/cobra
	$(GOGET) github.com/tarm/serial
	$(GOGET) go.bug.st/serial.v1
	$(GOGET) -u github.com/golang/protobuf/protoc-gen-go

# Cross compilation
#build-linux:
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v