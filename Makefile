# Use bash syntax
SHELL=/bin/bash

BUILD_TS:=$(shell date -u +"%Y-%m-%d_%H%M%S%Z")
BUILD_DIR:=./build
DIST_DIR:=./dist

APP_NAME:=my-go-template
APP_VERSION:=$(shell git describe --tags)

# subst meta data
PREFIX:=https://
SUFFIX:=.git
EMPTY:=

# Git parameters
GIT_BRANCH:=$(shell git rev-parse --abbrev-ref HEAD)
GIT_COMMIT:=$(shell git rev-parse HEAD)
GIT_REPO_DIR:=$(shell git rev-parse --show-toplevel)
GIT_REPO_URL:=$(shell git config --get remote.origin.url)
GIT_REPO:=$(subst $(PREFIX),$(EMPTY),$(subst $(SUFFIX),$(EMPTY),$(GIT_REPO_URL)))
#GIT_TAG:=$(shell git describe --abbrev=0 --tags)

# Go parameters
GOCMD=go
GOBINPATH=$(shell $(GOCMD) env GOPATH)/bin
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOENV=$(GOCMD) env
GOFMT=$(GOCMD) fmt
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
GOMOD=$(GOCMD) mod
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOTOOL=$(GOCMD) tool

GO_VERSION:=$(shell go version | sed -r 's/go version go(.*)\ .*/\1/')

GOBIN:=${GOPATH}/bin

GOFLAGS = -a
LDFLAGS = -s -w -X '$(GIT_REPO)/internal/version.AppVersion=$(APP_VERSION)' -X '$(GIT_REPO)/internal/version.Branch=$(GIT_BRANCH)' -X '$(GIT_REPO)/internal/version.BuildTime=$(BUILD_TS)' -X '$(GIT_REPO)/internal/version.Commit=$(GIT_COMMIT)' -X '$(GIT_REPO)/internal/version.GoVersion=$(GO_VERSION)'
#LDFLAGS = -s -w

# Tools
LINTER_REPORT = $(BUILD_DIR)/golangci-lint-$(BUILD_TS).out
COVERAGE_REPORT = $(BUILD_DIR)/unit-test-coverage-$(BUILD_TS)

# Rules
.PHONY: default
default: help

.PHONY: install
install:
	@echo "install"
	@./install-deps.sh

.PHONY: init
init:

.PHONY: clean
clean:
	@echo "clean"
	@rm -rf $(BUILD_DIR)
	@rm -rf $(DIST_DIR)
	@rm -f *.pprof
	@rm -f .DS_Store

.PHONY: goclean
goclean:
	@echo "goclean"
	@$(GOCLEAN) -cache -testcache -fuzzcache -x

.PHONY: godeepclean
godeepclean:
	@echo "godeepclean"
	@$(GOCLEAN) -cache -testcache -fuzzcache -modcache -x

.PHONY: $(BUILD_DIR)
$(BUILD_DIR):
	@echo "$(BUILD_DIR)"
	@mkdir -p $@

.PHONY: $(DIST_DIR)
$(DIST_DIR):
	@echo "$(DIST_DIR)"
	@mkdir -p $@

go.mod:
	@echo "go mod tidy"
	@$(GOMOD) tidy
	@echo "go mod verify"
	@$(GOMOD) verify
	@echo "go mod vendor"
	@$(GOMOD) vendor

go.sum: go.mod

.PHONY: fmt
fmt:
	@echo "go fmt"
	@$(GOFMT) ./...

.PHONY: prebuild
prebuild: init clean goclean $(BUILD_DIR) $(DIST_DIR) go.mod
	@echo "prebuild"
	$(GOCMD) version
	$(GOENV)

.PHONY: golangcilint
golangcilint: init $(BUILD_DIR)
	echo "Running golangci-lint"
	${GOPATH}/bin/golangci-lint --version
	${GOPATH}/bin/golangci-lint run --verbose --tests=true --timeout=1m --config .github/linters/.golangci.yml \
	  --issues-exit-code=0 --out-format=checkstyle > "$(LINTER_REPORT)"
	cat $(LINTER_REPORT)

.PHONY: lint
lint: golangcilint

.PHONY: gitleaks
gitleaks: init $(BUILD_DIR)
	@echo "Running gitleaks"
	gitleaks detect --config=.github/linters/.gitleaks.toml --source=. --redact --log-level=debug --report-format=json \
	  --report-path=$(BUILD_DIR)/gitleaks-$(BUILD_TS).out --verbose

.PHONY: unittest
unittest: init $(BUILD_DIR)
	$(GOENV)
	$(GOCMD) test -race -coverprofile="$(COVERAGE_REPORT).gcov" -covermode=atomic ./...
	cat "$(COVERAGE_REPORT).gcov"
	gcov2lcov -infile "$(COVERAGE_REPORT).gcov" -outfile "$(COVERAGE_REPORT).lcov"
	cat "$(COVERAGE_REPORT).lcov"
	$(GOCMD) tool cover -func="$(COVERAGE_REPORT).gcov"
#	$(GOCMD) tool cover -html="$(COVERAGE_REPORT).gcov"

.PHONY: tests
tests: unittest

.PHONY: gobuild
gobuild: $(DIST_DIR) prebuild lint gitleaks tests
	$(GOBUILD) $(GOFLAGS) -ldflags="$(LDFLAGS)" -o $(DIST_DIR)/$(APP_NAME) cmd/$(APP_NAME)/main.go

.PHONY: debug
debug: GOFLAGS += -x -v
debug: clean goclean gobuild

.PHONY: release
release: clean godeepclean gobuild

.PHONY: pre-commit
pre-commit: init
	pre-commit run --all-files

.PHONY: usage
usage:
	@echo "usage:"
	@echo "  make [command]"
	@echo "available commands:"
	@echo "  clean - clean up build artifacts"
	@echo "  goclean - call 'go clean -cache -testcache -fuzzcache -x'"
	@echo "  godeepclean - call 'go clean -cache -testcache -fuzzcache -modcache -x'"
	@echo "  debug - build debug version of binary"
	@echo "  help - show usage"
	@echo "  install - install latest build app dependancies (ie: golangci-lint, gcov2lcov)"
	@echo "  lint - run all linter checks"
	@echo "  release - build release version of binary"
	@echo "  tests - run all tests"
	@echo "  usage - show this information"

.PHONY: help
help: usage
