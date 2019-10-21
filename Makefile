PROJECT_NAME := "app"
PKG := "gitlab.com/marvincaspar/$(PROJECT_NAME)"
MAIN_FILE := "cmd/server/main.go"
PKG_LIST := $(shell go list ./... | grep -v /vendor/)
GO_FILES := $(shell find . -name '*.go' | grep -v /vendor/ | grep -v _test.go)
# https://github.com/golangci/awesome-go-linters
LINTERS := \
	github.com/rakyll/gotest \
	golang.org/x/lint/golint \
	honnef.co/go/tools/cmd/staticcheck

.PHONY: all init dep build clean test coverage coverhtml lint golint vet staticcheck

all: build

init: dep testdep ## Download dependencies and add git hooks
	find .git/hooks -type l -exec rm {} \;
	find githooks -type f -exec ln -sf ../../{} .git/hooks/ \;

lint: testdep ## Lint files
	@golint -set_exit_status ${PKG_LIST}

vet: testdep ## Checks correctness 
	@go vet ${PKG_LIST}

staticcheck: testdep ## Analyses code
	staticcheck ${PKG_LIST}

test: ## Run unit tests
	@go test -short ${PKG_LIST}

test-coverage: ## Run tests with coverage
	@go test -short -coverprofile cover.out -covermode=atomic ${PKG_LIST} 
	@cat cover.out >> coverage.txt

test-int: ## Run unit and integration tests
	@go test -short -tags=integration ${PKG_LIST}

coverage: ## Generate global code coverage report
	./scripts/coverage.sh;

coverhtml: ## Generate global code coverage report in HTML
	./scripts/coverage.sh html;

dep: ## Get dependencies
	@go mod tidy
	@go mod vendor

testdep: ## Get dev dependencies
	@go get -v $(LINTERS)

run:
	./bin/$(PROJECT_NAME)
	
build: dep ## Build the binary file
	@go build -i -o ./bin/$(PROJECT_NAME) ./$(MAIN_FILE)

clean: ## Remove previous build
	@rm -f ./bin

help: ## Display this help screen
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
