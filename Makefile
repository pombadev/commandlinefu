VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' -X 'main.revision=$(REVISION)'
GO ?= GO111MODULE=on go

.PHONY: build
build: main.go
	$(GO) build -ldflags "$(LDFLAGS)"

.PHONY: run
run: main.go
	$(GO) run -ldflags "$(LDFLAGS)" .
