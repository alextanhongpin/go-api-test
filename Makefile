include .env
export


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

info:
	@echo App Name: $(APP_NAME)
	@echo App Version: $(APP_VERSION)
	@echo Git Url: $(VCS_URL)
	@echo Git Hash: $(VCS_REF)
	@echo Build Date: $(BUILD_DATE)

run: wire
	@echo 'Starting the server'
	@go run main.go

install:
	@echo 'Installing external dependencies'
	@go get github.com/google/wire/cmd/wire

wire:
	@echo 'Generating dependencies injection using Wire'
	@wire ./...
