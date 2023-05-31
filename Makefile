GO ?= go
DOCKER ?= docker

.PHONY: test
test: generate-mocks
	$(GO) test --race ./...
.PHONY: generate-mocks
generate-mocks:
	$(DOCKER) run -v "$(PWD)":/app -w /app vektra/mockery
