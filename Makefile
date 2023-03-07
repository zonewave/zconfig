GO_FILES=`go list ./... | grep -v -E "mock|store|test|fake|cmd"`



lint: ## Lint Golang files
	@golint  ${GO_FILES}

.PHONY: test
test:
	@go test $(GO_FILES) -coverprofile .cover.txt
	@go tool cover -func .cover.txt

clean: ## Remove previous build
	@rm .cover.txt

up_pkg:
	@go get -u github.com/zonewave/pkgs@latest
	@go mod tidy