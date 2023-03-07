GO_FILES=`go list ./... | grep -v -E "mock|store|test|fake|cmd"`

#GOBIN = $(shell go env GOPATH)/bin
GOLINT =golint

MODULES = . ./tools



.PHONY: install
install:
	cd tools && go install golang.org/x/lint/golint

lint:
	@golint  -set_exit_status=1 ./...

golangci-lint:
	@golangci-lint run

.PHONY: test
test:
	go test -race -coverprofile=cover.out -coverpkg=./... ./...
	go tool cover -html=cover.out -o cover.html


clean: ## Remove previous build
	@rm cover.out
	@rm cover.html

up_pkg:
	@go get -u github.com/zonewave/pkgs@latest
	@go mod tidy