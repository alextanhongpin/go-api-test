include .env
export

include Makefile.*.mk

BUILD_DATE := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
# $ date -R
# Wed, 11 Jul 2018 10:25:26 +0800

# $ date -u +"%Y-%m-%dT%H:%M:%SZ"
# 2018-07-11T02:25:23Z

# VCS stands for Version Control System.
VCS_URL := $(shell git config --get remote.origin.url)

# $(shell git log -1 --pretty=%h)
VCS_REF := $(shell git rev-parse HEAD)


APP_VERSION := $(shell cat -s VERSION)
APP_NAME := $(shell basename `git rev-parse --show-toplevel`)


mockery := go run github.com/vektra/mockery/cmd/mockery

all: run

info:
	@echo App Name: $(APP_NAME)
	@echo App Version: $(APP_VERSION)
	@echo Git Url: $(VCS_URL)
	@echo Git Hash: $(VCS_REF)
	@echo Build Date: $(BUILD_DATE)


run: wire
	@echo 'Starting the server'
	@go run {main,wire_gen}.go


install:
	@echo 'Installing external dependencies'
	@go get github.com/google/wire/cmd/wire
	@go install github.com/vektra/mockery/cmd/mockery@v1.1.2
	@go get github.com/vektra/mockery/...


test:
	@echo "$(OK_COLOR)Running test coverage$(NO_COLOR)"
	@go test -v -failfast -cover -coverprofile=cover.out ./...
	@go tool cover -html=cover.out


wire: # NOTE: Running go generate will also work
	@echo "$(OK_COLOR)Generating dependencies injection using Wire$(NO_COLOR)"
	@wire ./...


mock: # Generates mocks for the given interface name.
	@echo "$(OK_COLOR)Generating mock for $(name)$(NO_COLOR)"
	@$(mockery) -name $(name) -recursive -case underscore
