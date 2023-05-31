GO ?= go
DOCKER ?= docker

.PHONY: test
test:
	$(GO) test --race ./...
.PHONY: list
list:
	$(GO) list ./...
.PHONY: generate-mocks
generate-mocks:
	$(DOCKER) run -v "$(PWD)":/app -w /app vektra/mockery