VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "0.0.1")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"

.PHONY: update-vendor
update-vendor:
	go mod tidy
	go mod vendor

.PHONY: clean
clean:
	rm -rf dist/*

.PHONY: start
start:  ## build the binary
	CGO_ENABLED=0 go run ${LDFLAGS} main.go

.PHONY: build
build:  ## build the binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a

.PHONY: version
version: ## display the version of the binary
	@echo $(VERSION)

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: format
format: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: test
test: ## run unit tests
	go test -v ./...
		