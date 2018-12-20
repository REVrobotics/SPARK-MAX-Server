# Basic go commands
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get

# Binary names
BINARY_NAME_WINDOWS=sparkmax.exe
BINARY_UNIX=sparkmax.out

# Versioning, increment every time called
BUILD_NUMBER_FILE := build-number.txt
BUILD_DATE := `date -u +.%Y%m%d.%H%M%S`
GIT_COMMIT := $(shell git describe --always --long --dirty)

CURRENT_DIR=$(shell pwd)
PROJECT_PATH=${GOPATH}/src/github.com/REVrobotics/SPARK-MAX-Server/

ifndef GOPATH
  $(error GOPATH is undefined, must have a working go distribution with GOPATH environtment variable set)
endif

ifeq ($(OS),Windows_NT)
	BINARY_NAME=${BINARY_NAME_WINDOWS}
else
	BINARY_NAME=${BINARY_UNIX}
endif

export CGO_CFLAGS=-I${PROJECT_PATH} -g -O2
export CGO_LDFLAGS=-L${PROJECT_PATH} -g -O2

$(info ${CGO_CFLAGS})

all: build
build: 
	@echo $$(($$(cat $(BUILD_NUMBER_FILE)) + 1)) > $(BUILD_NUMBER_FILE)
	protoc -I./sparkmax --go_out=./sparkmax ./sparkmax/SPARK-MAX*.proto
	$(GOBUILD) -o $(BINARY_NAME) -v -ldflags "-s -w -X main.BuildNumber=`cat ${BUILD_NUMBER_FILE}` -X 'main.BuildDate=`date +%Y-%m-%d%_I:%M:%S`' -X main.BuildCommit=${GIT_COMMIT}"

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
	$(GOGET) github.com/willtoth/go-dfuse
	$(GOGET) gopkg.in/cheggaaa/pb.v1

# Cross compilation
#build-linux:
#	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v