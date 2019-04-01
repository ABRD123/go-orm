SHELL = /bin/bash

CMD_DIRS := \
	./cmd/dbinit\
	./cmd/user

SRC_DIRS := $(shell \
	find . -name "*.go" -not -path "./vendor/*" | \
	xargs -I {} dirname {}  | \
	uniq)

TOOLS := \
	github.com/golang/dep/cmd/dep; \
	golang.org/x/lint/golint; \
	golang.org/x/tools/cmd/goimports; \
	github.com/stretchr/testify;

GETTOOLS := $(foreach TOOL,$(TOOLS),go get -u $(TOOL))

default: tools deps fmt vet lint build

all: tools deps fmt vet lint build

prod:
	go install -v $(CMD_DIRS)

install: tools deps fmt vet lint
	go install -v $(CMD_DIRS)

tools:
	$(GETTOOLS)

deps:
	dep ensure

fmt: $(SRC_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		goimports -w -local "github.com/go-orm/" *.go ; \
		popd > /dev/null ; \
	done;

vet: $(SRC_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go vet ; \
		popd > /dev/null ; \
	done;

lint: $(SRC_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		golint -set_exit_status ; \
		popd > /dev/null ; \
	done;

build: $(CMD_DIRS)
	@for dir in $^; do \
		pushd ./$$dir > /dev/null ; \
		go build -v ; \
		popd > /dev/null ; \
	done;

clean:
	go clean $(CMD_DIRS)

uninstall:
	go clean -i $(CMD_DIRS)

.PHONY: check prod all tools dep fmt vet lint test install build clean
