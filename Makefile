.DEFAULT_GOAL := help

SHELL := /bin/bash

GRAFANA_CH_DS_VERSION=4.5.1
GRAFANA_CH_DS_ARCH=linux_arm64
GRAFANA_CH_DS_URL="https://github.com/grafana/clickhouse-datasource/releases/download/v$(GRAFANA_CH_DS_VERSION)/grafana-clickhouse-datasource-$(GRAFANA_CH_DS_VERSION).$(GRAFANA_CH_DS_ARCH).zip"

# constant variables
PROJECT_NAME	= c24-expences
BINARY_NAME	= c24-expences
GIT_COMMIT	= $(shell git rev-parse HEAD)
BINARY_TAR_DIR	= $(BINARY_NAME)-$(GIT_COMMIT)
BINARY_TAR_FILE	= $(BINARY_TAR_DIR).tar.gz
BUILD_VERSION	= $(shell cat VERSION.txt)
BUILD_DATE	= $(shell date -u '+%Y-%m-%d_%H:%M:%S')

# LDFLAGS=-ldflags "-X=main.Version=$(VERSION) -X=main.Build=$(BUILD)"

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

# golangci-lint config
golangci_lint_version=v1.64.8
vols=-v `pwd`:/app -w /app
run_lint=docker run --rm $(vols) golangci/golangci-lint:$(golangci_lint_version)

.PHONY: lint fmt test help

## help: prints this help message
help:
	@echo "Usage:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

## fmt: format the code
fmt:
	@gofmt -l -w $(SRC)

## lint: run linters
lint:
	@printf "$(OK_COLOR)==> Running golang-ci-linter via Docker$(NO_COLOR)\n"
	@$(run_lint) golangci-lint run --timeout=5m --verbose
#

## build: test compile the binary for the local machine
build:
	@echo 'compiling binary...'
	@cd cmd/ && CGO_ENABLED=0 go build -o ../$(BINARY_NAME)


## setup: install dependencies and setup the project
setup:
	@printf "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)\n"
	@go mod download
	@mkdir data_dir dist
	@wget -O dist/ch.zip $(GRAFANA_CH_DS_URL) && unzip dist/ch.zip -d dist


## test: run tests with coverage
test:
	@printf "$(OK_COLOR)==> Running tests$(NO_COLOR)\n"
	@go test -v -count=1 -covermode=atomic -coverpkg=./... -coverprofile=coverage.txt ./...
	@go tool cover -func coverage.txt
